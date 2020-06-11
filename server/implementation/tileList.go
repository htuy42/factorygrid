package implementation

import (
	proto "server/protos"
	"sync"
	"time"
)

// Just a linked list of tileListener, very basic implementation.
// can be safely read by any number of threads, whenever. Its always in a safe state, but updates aren't guaranteed
// at any specific time. It synchronizes its own writing. Right now it uses a mtx, this could easily be modified to
// stream updates onto a chan and allow only a single goroutine to modify it, but I don't think this is necessary rn
type listenerList struct {
	tileListHead *tileListLink
	tileListTail *tileListLink
	mtx          sync.Mutex
}

type tileListLink struct {
	listener *tileListener
	next     *tileListLink
	prev     *tileListLink
	owner    *listenerList
}

// Create a new tileListener for the given server, and add it to t. Also starts all of the background processes needed
// for the listener
func (t *listenerList) startListener(server proto.FactoryService_RequestViewStreamServer, f *factoryInner, batchIntervalMs int64) chan bool {
	listener := &tileListener{}
	listener.changeChan = make(chan *tile, 1000)
	listener.changedTiles = make(map[*tile]bool)
	listener.f = f
	listener.server = server
	listener.done = false
	listener.doneChan = make(chan bool)
	listener.sendTimeout = make(chan bool)
	listener.batchIntervalMs = batchIntervalMs

	// This line exposes the listener to the list
	listener.ownLink = t.add(listener)
	go listener.serverReceive()
	go listener.dirtyTileHandle()
	go listener.changeTimer()
	return listener.doneChan
}

// create a listener list
func makeListenerList() *listenerList {
	res := listenerList{}
	// not strictly necessary but done for clarity
	res.tileListTail = nil
	res.tileListHead = nil
	return &res
}

// Add a tileListener to the list. Returns the link the tileListener is now attached to.
func (t *listenerList) add(listener *tileListener) *tileListLink {
	t.mtx.Lock()
	defer t.mtx.Unlock()
	link := &tileListLink{listener: listener, prev: t.tileListTail, next: nil, owner: t}
	t.tileListTail = link
	if t.tileListHead == nil {
		t.tileListHead = link
	} else {
		link.prev.next = link
	}
	return link
}

// Remove the given link from the list.
func (t *listenerList) remove(link *tileListLink) {
	t.mtx.Lock()
	defer t.mtx.Unlock()
	if t.tileListHead == link {
		t.tileListHead = link.next
	}
	if t.tileListTail == link {
		t.tileListTail = link.prev
	}
	if link.prev != nil {
		link.prev.next = link.next
	}
	if link.next != nil {
		link.next.prev = link.prev
	}
}

// fan out a tile update to all of the tileListeners linked to this list.
func (t *listenerList) sendChannelChanges(tile *tile) {
	listItem := t.tileListHead
	for listItem != nil {
		select {
		case listItem.listener.changeChan <- tile:
		default:
		}
		listItem = listItem.next
	}
}

// receives requests from a client, and sends it updated views of the factory.
type tileListener struct {
	// server managed directly by this listener.
	server          proto.FactoryService_RequestViewStreamServer
	// The rectangle the client this listener is serving is currently interested in. Used to determine whether or not
	// to send changed tiles (ie only tiles that fall within this rect will be forwarded
	currentRect     *proto.Rectangle
	// the listener gets to hear about all the changed tiles on this channel
	changeChan      chan *tile
	// used for timing the batching of viewResponses.
	sendTimeout     chan bool
	// A set of tiles that have been changed since we last sent a batch update
	changedTiles    map[*tile]bool
	// The listeners own list in the listenerList. Used for removing it when it is destroyed.
	ownLink         *tileListLink
	// Whether or not the listener is done. When this becomes true, the background processes it runs will be shut down.
	done            bool
	// a reference to the factory that created this listener, in order to get access to the tiles for view square updates
	f               *factoryInner
	// true is sent on this when the listener finishes so the server can be shutdown.
	doneChan        chan bool
	// How many ms should be waited before each batch update is sent.
	batchIntervalMs int64
}


// Destroy the listener. Marks it as done, removes its link from the list, and sends true on its doneChan (so it server
// will be destroyed)
func (l *tileListener) destroy() {
	l.ownLink.owner.remove(l.ownLink)
	l.done = true
	l.doneChan <- true
}

// Send a batch of tiles for new tiles the user is viewing when they move their display.
func (l *tileListener) batchNewViewValues(cellRect *proto.Rectangle, oldRect *proto.Rectangle) {
	sendTo := cellRect
	resp, err := l.f.RequestViewSquares(sendTo, oldRect)
	if err != nil {
		l.destroy()
	} else {
		// Don't sent the update if there aren't actually any new tiles (ie the user scrolled one cell to the right,
		// but not into a new tile.
		if len(resp.SubViews) > 0{
			err := l.server.Send(resp)
			if err != nil {
				l.destroy()
			}
		}

	}
}

// Receive updates from the client to the view rect it is interested in.
func (l *tileListener) serverReceive() {
	for !l.done {
		request, err := l.server.Recv()
		if err != nil {
			l.destroy()
			return
		}
		rect := request.FullView
		l.batchNewViewValues(rect, request.OldView)

		// adjust rect to tile space from cell space
		rect.StartX /= l.f.tileSize32
		rect.StartY /= l.f.tileSize32

		// YUCK this is 2 lines with ternary
		if rect.Width%l.f.tileSize32 != 0 {
			rect.Width /= l.f.tileSize32
			rect.Width += 1
		} else {
			rect.Width /= l.f.tileSize32
		}
		if rect.Height%l.f.tileSize32 != 0 {
			rect.Height /= l.f.tileSize32
			rect.Height += 1
		} else {
			rect.Height /= l.f.tileSize32
		}
		l.currentRect = rect
	}
}

// Send a batch update of all the tiles that have changed since our last update for this listener.
// todo its possible we ought to be grouping these batch sends up so that all the listeners interested in a set of
// tiles require only a single acquisition of the tile locks.
func (l *tileListener) sendBatchedUpdates() {
	response := proto.ScreenResponse{}
	response.SubViews = make([]*proto.ViewResponse, len(l.changedTiles))
	ind := 0
	// Acquire the locks for every single one of the tiles we are viewing. Starts goroutines to do this so we are
	// immediately waiting for all of them. Also create the response: note that right now we are really just gathering
	// references for it, which we don't need locks to do. We need the locks when we are sending because at that point
	// the bytes the references point to are copied to wire, and they can't change during that process.

	// todo consider: all the changes we make *could* be made safely if they didn't change the *size* of the tiles.
	// In order to do this, we could require that every tile have either only a single entity, or at most some fixed
	// number of entities, and have the viewResponses always send the full list. This just requires creating a 0
	// value for entities that the client knows to ignore, and slightly increasing the bite size of messages. This would
	// eliminate the slightly kludgy locking that we currently "get" to do.
	lockedChan := make(chan bool)
	for element := range l.changedTiles {
		go func(t *tile) {
			t.lastViewResponseLock.Lock()
			lockedChan <- true
		}(element)
		response.SubViews[ind] = element.lastViewResponse
		ind++
	}
	// wait till we have all the locks
	for i := 0; i < len(l.changedTiles); i++ {
		<-lockedChan
	}
	// Once we have all the locks, send the response on the wire.
	err := l.server.Send(&response)
	if err != nil {
		l.destroy()
	}
	// unlock all the locks. We don't need to go this since unlocking isn't blocking.
	for element := range l.changedTiles {
		element.lastViewResponseLock.Unlock()
	}
	l.changedTiles = make(map[*tile]bool)
}

// Receive dirty tile updates, and add them to our batch. When the batch timer goes off, send an update batch.
func (l *tileListener) dirtyTileHandle() {
	for !l.done {
		select {
		case _ = <-l.sendTimeout:
			l.sendBatchedUpdates()
			// pass the timeout token back to changeTimer. This way the batch update doesn't start counting down again
			// until the batch has *finished*
			// todo consider: whether this is how we actually want it or not. There are arguments either way: with this
			// implementation, if the actual batch locking + sending becomes a bottleneck, we still get some down time
			// between batches, where with the other way if a batch send takes batchIntervalMs the next one will happen
			// immediately when the first one finishes.
			// On the other hand, with the other way, batchIntervalMs actually represents the time between successful
			// batch sends, which is to say that 1000 / batchIntervalMs updates will be sent every second unless the
			// server bottlenecks, which isn't true with this implementation
			l.sendTimeout <- true
		case change := <-l.changeChan:
			cX := int32(change.x)
			cY := int32(change.y)
			// grab a reference to the rect we are using so that it doesn't change mid-handling
			referenceRect := l.currentRect
			// only send the change to the client if it is interested in this tile
			if cX >= referenceRect.StartX && cX < referenceRect.StartX+referenceRect.Width && cY >= referenceRect.StartY && cY < referenceRect.StartY+referenceRect.Height {
				l.changedTiles[change] = true
			}
		}
	}
}

// every batchIntervalMs, trigger a new batch. Blocks on the send,
func (l *tileListener) changeTimer() {
	for !l.done {
		// sleep, then send the timeout token to the batch sender. See the note in dirtyTileHandle for what this is and
		// some considerations on its implications.
		time.Sleep(time.Duration(l.batchIntervalMs) * time.Millisecond)
		l.sendTimeout <- true
		<- l.sendTimeout
	}
}
