package implementation

import (
	"server/config"
	proto "server/protos"
)


// Factory is a confusing name: a more accurate one would be "world" or "grid controller," but this project was
// loosely motivated by creating a factorio "clone," so Factory it is.

// Interface providing the necessary functionality to serve the grpc requests for a FactoryService
type Factory interface {
	// Begin the internal update functions of the factory. Other functions shouldn't be called until this one has.
	StartRunning()
	// Give the responses for a given set of view squares. oldRect denotes the last view rectangle this client knows
	// about. The client will get a ScreenResponse for rectangle - oldRect, to avoid sending tiles the client already
	// knows about
	RequestViewSquares(rectangle *proto.Rectangle, oldRect *proto.Rectangle) (*proto.ScreenResponse, error)
	// Send an interaction and have it handled by the tile it designates as its target
	SendInteraction(interaction *proto.Interaction)
	// Instruct the factory to begin serving a RequestViewStream. The returned channel outputs a boolean when the
	// stream is done, before which point the server should not be exited.
	AddViewStream(server proto.FactoryService_RequestViewStreamServer) chan bool
}

// Implementation of Factory
type factoryInner struct {
	// The size of a given cell in tiles, as well as the number of tiles in the x and y direction for the whole world
	tileSize, worldWidthTiles, worldHeightTiles int
	// Currently unused, would provide a global sync rate for all the tiles so that they all ticked in more or less
	// lock step
	tileSync bool
	// int32 version of tileSize
	tileSize32 int32
	// All of the tiles. Note this is in y,x order, which is to say accessing tile (x,y) is accomplished by tiles[y][x]
	tiles [][]*tile
	// When a tile is changed and updates need to be sent to clients looking at it, it will send itself on this channel
	changedTilesChan chan *tile
	// List keeping track of the view streams we are currently serving
	tileListeners *listenerList
	// Whether or not the factory is currently running. Setting this to false will terminate the factory.
	running bool
	// When an entity moves from one tile to another, the tile will dump the entity to this. The factory will
	// then send the entity to the appropriate new host tile of the entity
	movingEntities chan *entity
	// Config accessor for the factory
	conf config.ConfigProvider
}


func (f *factoryInner) AddViewStream(server proto.FactoryService_RequestViewStreamServer) chan bool {
	return f.tileListeners.startListener(server,f, f.conf.GetConfigI64("world-configs","view-batch-interval-ms"))
}

func (f *factoryInner) SendInteraction(interaction *proto.Interaction) {
	tileX := interaction.X / f.tileSize32
	tileY := interaction.Y / f.tileSize32
	// As a rule, when an invalid coordinate appears we just drop the message
	if tileX >= 0 && tileX < int32(f.worldWidthTiles) && tileY >= 0 && tileY < int32(f.worldHeightTiles){
		interaction.X %= f.tileSize32
		interaction.Y %= f.tileSize32
		f.tiles[tileY][tileX].interactions <- interaction
	}
}

func (f *factoryInner) RequestViewSquares(rectangle *proto.Rectangle, oldRect *proto.Rectangle) (*proto.ScreenResponse, error) {
	var subviews []*proto.ViewResponse
	startSquareX := rectangle.StartX / f.tileSize32
	startSquareY := rectangle.StartY / f.tileSize32
	endSquareX := (rectangle.StartX + rectangle.Width) / f.tileSize32
	endSquareY := (rectangle.StartY + rectangle.Height) / f.tileSize32
	oldStartX := oldRect.StartX / f.tileSize32
	oldStartY := oldRect.StartY / f.tileSize32
	oldEndX := (oldRect.StartX + oldRect.Width) / f.tileSize32
	oldEndY := (oldRect.StartY + oldRect.Height) / f.tileSize32
	if (rectangle.StartX + rectangle.Width) % f.tileSize32 != 0{
		endSquareX += 1
	}
	if (rectangle.StartY + rectangle.Height) % f.tileSize32 != 0{
		endSquareY += 1
	}
	if (oldRect.StartX + oldRect.Width) % f.tileSize32 != 0{
		oldEndX += 1
	}
	if (oldRect.StartY + oldRect.Height) % f.tileSize32 != 0{
		oldEndY += 1
	}
	for x := startSquareX; x <= endSquareX; x++{
		for y := startSquareY; y <= endSquareY; y++{
			// skip if its within old rect!
			if x > oldStartX && x <= oldEndX && y > oldStartY && y <= oldEndY{
				continue
			}
			if x < 0 || y < 0 || x >= int32(f.worldWidthTiles) || y >= int32(f.worldHeightTiles){
				continue
			}
			subviews = append(subviews, f.tiles[y][x].lastViewResponse)
		}
	}
	return &proto.ScreenResponse{SubViews: subviews}, nil
}

// Create a new factory. In practice returns a factoryInner, for the moment
func MakeFactory(tileSize, worldWidthTiles, worldHeightTiles int, tileSync bool, msBetweenTicks int64, conf config.ConfigProvider) Factory {
	res := factoryInner{
		tileSize:         tileSize,
		worldWidthTiles:  worldWidthTiles,
		worldHeightTiles: worldHeightTiles,
		tileSize32:       int32(tileSize),
		tiles:            make([][]*tile,worldHeightTiles),
		running:          true,
		tileSync:         tileSync,
		changedTilesChan: make(chan *tile, 1000),
		tileListeners:    makeListenerList(),
		movingEntities:   make(chan *entity, 1000),
		conf:             conf,
	}
	for i := 0; i < len(res.tiles); i++ {
	    res.tiles[i] = make([]*tile,worldWidthTiles)
	    for j := 0; j < len(res.tiles[i]); j++ {
	        res.tiles[i][j] = makeTile(tileSync,res.changedTilesChan,j,i,tileSize,msBetweenTicks,res.movingEntities)
		}
	}
	return &res
}

// If tile sync is on, sends a sync token, then do not send a new one out to any of them until they have all sent it
// back. Tiles currently ignore sync tokens, but if this feature is wanted that can easily be changed.
func (f *factoryInner) beginTileSync(){
	for f.running{
		for i := 0; i < len(f.tiles); i++ {
		    cur := f.tiles[i]
		    for j := 0; j < len(cur); j++ {
		        tile := cur[i]
		        tile.syncControl <- true
		    }
		}
		for i := 0; i < len(f.tiles); i++ {
			cur := f.tiles[i]
			for j := 0; j < len(cur); j++ {
				tile := cur[i]
				<- tile.syncControl
			}
		}
	}
}

// Repeatedly pole the channel of entities that need to be moved, and forward entities to the tile that now owns them
func (f *factoryInner) transferEntities(){
	for f.running{
		movedEntity := <- f.movingEntities
		targetTileX := movedEntity.worldX / f.tileSize32
		targetTileY := movedEntity.worldY / f.tileSize32
		// If it falls off the edge, it gets thrown out.
		if int(targetTileX) < f.worldWidthTiles &&  movedEntity.worldX > 0{
			if int(targetTileY) < f.worldHeightTiles && movedEntity.worldY > 0{
				f.tiles[targetTileY][targetTileX].entitiesIn <- movedEntity
			}
		}
	}
}

// Pole tiles from the changedTilesChan and send them onto the tileListenerList
func (f *factoryInner) beginTileCaching(){
	for f.running{
		changedTile := <- f.changedTilesChan
		f.tileListeners.sendChannelChanges(changedTile)
	}
}

// Begin tileCaching and entityTransfer, as well as tileSync if that feature is on. Essentially run all of the
// background features the factory requires
func (f *factoryInner) StartRunning() {
	if f.tileSync{
		go f.beginTileSync()
	}
	go f.beginTileCaching()
	go f.transferEntities()
}
