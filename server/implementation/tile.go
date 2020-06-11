package implementation

import (
	proto "server/protos"
	"sync"
	"time"
)

// A single cell within a tile
type cell struct {
	// The typeId of the cell
	typeId int32
	// All of the entities in the cell. This is just a set: the bool is irrelevant
	entities map[*Entity]bool
	// Whether or not this cell changed since the last time a viewResponse was created for the cells host tile, and
	// therefore whether or not it needs to be updated in the next viewResponse
	dirty bool
}

// A single tile within the factory world
type tile struct {
	// All of the cells that make up the tile
	internalContents [][]cell
	// All of the entities that live in any of the cells. They are also present in the individual cell Entity maps, but
	// having a single global map of them is convenient
	allEntities map[*Entity]bool
	// syncControl channel on which sync tokens could be received if syncControl were being used. Currently ignored
	// since this features isn't in use
	syncControl chan bool
	// The tile sends a pointer to itself on this channel whenever it updates its view response
	changeChannel chan *tile
	// The x and y location of the tile within the broader grid. These are in tile, not cell coordinates. Also, the size
	// of the tile in cells: the tile is a square, so this is both width and height
	x, y, size int
	// The x and y location of the tile within the grid in terms of cell coordinates.
	xCell, yCell int32
	// int32 of the size to avoid some of the casting that would otherwise make me feel unhappy about golang
	size32 int32
	// Interactions for this tile are sent into this channel by the factory
	interactions chan *proto.Interaction
	// Entities that are being forwarded to the tile come in on this
	entitiesIn chan *Entity
	// Entities that have been moved out of the tile can be sent out on this, and will be forwarded to their new
	// host tiles entitiesIn channel
	entitiesOut chan *Entity
	// Whether or not syncControl is used. For now, always false
	waitsForSync bool
	// The lastViewResponse this tile generated. Repeatedly updated when the tile changes its view.
	lastViewResponse *proto.ViewResponse
	// The amount of time between ticks. Whenever the tile ticks, if it takes less than this amount of time it will
	// then sleep until at least this many ms have elapsed since it started its tick
	msBetweenTicks int64
	// Whether the entire tile is dirty. If this is false, it means that no cell in the tile is dirty
	dirty bool
	// The current generation of the tile, ie the number of ticks that have occurred since it was created
	generation int
	// A lock that the tile holds when it is changing the lastViewResponse, since that value is also read by tileList
	// when it sends the view response to clients.
	lastViewResponseLock sync.Mutex
	// A lock that lets the tile synchronize self-modification during ticks with viewResponse refreshing.
	modLock sync.Mutex
	// Whether or not a view response refresh is currently happening. Prevents more than one refresh from being queued
	// up, since there is nothing wrong with doing several modifications and then only refreshing the view once.
	refreshing     bool
	interactionMap map[string]Interaction
}

// Refresh lastViewResponse for changes that have happened due to ticks since it was last refreshed
func (t *tile) refreshViewResponseFromCells() {
	// Get the response lock first, which gives us permission to change the viewResponse.
	t.lastViewResponseLock.Lock()
	// Only grab the modlock once we have the refresh lock. We don't really mind waiting to refresh, since that means
	// something is reading us. We do not want to get stuck not being able to mod while someone else is sending us on
	// wire, though.
	t.modLock.Lock()
	defer t.lastViewResponseLock.Unlock()
	defer t.modLock.Unlock()
	// For each dirty cell, update its values in lastViewResponse
	for x := 0; x < t.size; x++ {
		for y := 0; y < t.size; y++ {
			reference := &t.internalContents[y][x]
			if reference.dirty {
				t.lastViewResponse.Tiles[y*t.size+x].TileTypeId = reference.typeId
				t.lastViewResponse.Tiles[y*t.size+x].Entities = make([]*proto.Entity, len(reference.entities))
				ind := 0
				for e := range reference.entities {
					t.lastViewResponse.Tiles[y*t.size+x].Entities[ind] = e.internal
					ind++
				}
				reference.dirty = false
			}
		}
	}

	// We are no longer dirty, and refreshing is done
	t.dirty = false
	t.refreshing = false
	// We have now changed lastViewResponse, so inform the factory that we can be resent.
	t.changeChannel <- t
}

// Setup the tile. Also starts its ticking
func makeTile(waitForSync bool, outChan chan *tile, x, y, size int, msBetweenTicks int64, entitiesOut chan *Entity,
	interactions map[string]Interaction) *tile {
	res := tile{}
	res.allEntities = make(map[*Entity]bool)
	res.syncControl = make(chan bool, 2)
	res.changeChannel = outChan
	res.interactions = make(chan *proto.Interaction, 1000)
	res.interactionMap = interactions
	res.waitsForSync = waitForSync
	res.x = x
	res.y = y
	res.xCell = int32(x * size)
	res.yCell = int32(y * size)
	res.size = size
	res.entitiesOut = entitiesOut
	res.entitiesIn = make(chan *Entity, 100)
	res.refreshing = false
	res.size32 = int32(res.size)
	res.setInitialContents()
	res.msBetweenTicks = msBetweenTicks
	go res.startTicking()
	return &res
}

// Create the initial contents of the tile. Currently just sets every cell to typeId 0 with no entities.
func (t *tile) setInitialContents() {
	t.internalContents = make([][]cell, t.size)
	for y := 0; y < t.size; y++ {
		t.internalContents[y] = make([]cell, t.size)
		for x := 0; x < t.size; x++ {
			t.internalContents[y][x] = cell{}
			t.internalContents[y][x].dirty = true
			t.internalContents[y][x].typeId = int32(0)
			t.internalContents[y][x].entities = make(map[*Entity]bool)
		}
	}
	t.lastViewResponse = &proto.ViewResponse{}
	t.lastViewResponse.ViewOf = &proto.Rectangle{
		StartX: int32(t.x),
		StartY: int32(t.y),
		Width:  int32(t.size),
		Height: int32(t.size),
	}
	t.lastViewResponse.Tiles = make([]*proto.Tile, t.size*t.size)
	for i := 0; i < t.size*t.size; i++ {
		t.lastViewResponse.Tiles[i] = &proto.Tile{}
	}
	t.refreshViewResponseFromCells()
}

// Run a single Interaction in the tile
func (t *tile) runInteraction(interaction *proto.Interaction) {
	if interaction.X < 0 || interaction.Y < 0 {
		return
	}
	affectedCell := &t.internalContents[interaction.Y][interaction.X]
	// Since the large majority of interactions dirty the tile, we default to marking it dirty and undirty it manually
	// for interactions that *do not* actually need it dirty. This is just for code brevity.
	i, ok := t.interactionMap[interaction.InteractionChar]
	if ok {
		if i.trigger(t, affectedCell, interaction.X, interaction.Y, interaction.X+t.xCell, interaction.Y+t.yCell) {
			affectedCell.dirty = true
			t.dirty = true
		}
	}
}

// Add an Entity, either a new one that was just spawned or one that was forwarded from another tile
func (t *tile) addEntity(e *Entity) {
	destX := e.worldX % t.size32
	destY := e.worldY % t.size32
	t.internalContents[destY][destX].entities[e] = true
	t.internalContents[destY][destX].dirty = true
	t.allEntities[e] = true
	t.dirty = true
}

// Pull all interactions and forwarded entities until there are none left, and add / run them
func (t *tile) runInteractionsAndEntityMoves() {
	running := true
	for running {
		select {
		case i := <-t.interactions:
			t.runInteraction(i)
		case e := <-t.entitiesIn:
			t.addEntity(e)
		default:
			running = false
		}
	}
}

// Blocks until either a new Interaction or a new Entity to come into the cell, since there's no need to tick while
// we have none of either
func (t *tile) blockForChange() {
	select {
	case e := <-t.entitiesIn:
		t.entitiesIn <- e
	case i := <-t.interactions:
		t.interactions <- i
	}
}

// Move an Entity that has updated its own world location in its last tick.
// returns whether or not the Entity moved to an entirely new block.
func (t *tile) moveEnt(ent *Entity, startLocalX, startLocalY int32) bool {
	t.dirty = true
	t.internalContents[startLocalY][startLocalX].dirty = true
	delete(t.internalContents[startLocalY][startLocalX].entities, ent)
	if ent.worldX/t.size32 == int32(t.x) && ent.worldY/t.size32 == int32(t.y) {
		t.internalContents[ent.worldY%t.size32][ent.worldX%t.size32].entities[ent] = true
		t.internalContents[ent.worldY%t.size32][ent.worldX%t.size32].dirty = true
		return false
	} else {
		t.entitiesOut <- ent
		return true
	}
}

// Tick all entities. Marks us dirty if any new entities are created, any entities are deleted, or any Entity moves.
// Otherwise, there may have been Entity state updates, but there are no changes the client can see, so we don't need
// to be dirty.
func (t *tile) tickEntities() {
	// TODO: entities will  theoretically appear in the block they are leaving and the block they are entering briefly:
	// they are moved, and then the new tile finds about them. At that point it might send off an update with them
	// added, before this block finishes ticking and sends off an update with them removed. They won't actually be in
	// both places logically from the servers perspective / we aren't really duplicating entities, but the client may
	// view both for a moment.  On the plus side, its "eventually consistent" XD
	var newEntities []*Entity
	var deletedEntities []*Entity
	for ent := range t.allEntities {
		curX := ent.worldX
		curY := ent.worldY
		ent.tick(t)
		if ent.dead {
			deletedEntities = append(deletedEntities, ent)
		}
		newEntities = append(newEntities, ent.children...)
		ent.children = []*Entity{}
		if ent.worldX != curX || ent.worldY != curY && !ent.dead {
			removed := t.moveEnt(ent, curX%t.size32, curY%t.size32)
			if removed {
				deletedEntities = append(deletedEntities, ent)
			}
		}
	}
	if len(deletedEntities) != 0 {
		t.dirty = true
	}
	for _, ent := range deletedEntities {
		delete(t.allEntities, ent)
	}
	for _, ent := range newEntities {
		// If it doesn't belong in this block, dump it to wherever it does belong.
		if ent.worldX/t.size32 == int32(t.x) && ent.worldY/t.size32 == int32(t.y) {
			t.addEntity(ent)
		} else {
			t.entitiesOut <- ent
		}
	}
}

func (t *tile) startTicking() {
	for true {
		// If we have any entities we don't want to block, but if we don't we don't need to do anything till further
		// prompting
		if len(t.allEntities) == 0 {
			t.blockForChange()
		}
		// For syncing to t.msBetweenTicks
		tickTime := time.Now()
		// Lock the mod lock in order to actually change contents
		t.modLock.Lock()
		t.tickEntities()
		t.runInteractionsAndEntityMoves()
		t.modLock.Unlock()
		// do actual ticking stuff
		// Only refresh if we are dirty and not already refreshing.
		if t.dirty {
			if !t.refreshing {
				t.refreshing = true
				// This holds the modlock. We are happy to be able to do it while we are otherwise waiting for stuff to
				// happen, so we go it.
				go t.refreshViewResponseFromCells()
			}
		}
		t.generation += 1
		elapsed := time.Now().Sub(tickTime)
		if elapsed.Milliseconds() < t.msBetweenTicks {
			time.Sleep(time.Duration(t.msBetweenTicks-elapsed.Milliseconds()) * time.Millisecond)
		}
	}
}
