package implementation

import proto "server/protos"

// An Entity within the grid
type Entity struct{
	// The wire representation of the Entity. The client doesn't need state information about the Entity, just its
	// type and location so it knows how to draw it.
	internal *proto.Entity
	// The location of the Entity in world coordinates.
	worldX, worldY int32
	// During a given tick, all the new entities this one has created. At the end of a tick, this list will be
	// emptied and all the entities in it added to the world.
	children []*Entity
	// Whether or not the Entity is dead. If this is set to true, the Entity will be removed.
	dead bool
	// The ticker that defines specifically what type of Entity this is.
	ticker EntityTicker
}

// Provides the actual identity of a type of Entity and its control logic.
type EntityTicker interface{
	// Get the entities type id
	getTypeId() int32
	// Update the Entity for a new world tick. Can read the tile, but should generally not modify it. There aren't
	// any specific controls in place to enforce this, so its easy to break things if you want to.
	// todo put Entity in a different package and obscure tile private members behind accessor functions so they can't
	// be messed up as easily
	tick(e *Entity, t *tile)
}

// Just call the ticker, doesn't need to do anything else.
func (e *Entity) tick(t *tile){
	e.ticker.tick(e,t)
}

// Create an Entity at the given location with the given ticker
func makeEntity(startX, startY int32, ticker EntityTicker) *Entity {
	res := Entity{
		internal: &proto.Entity{TypeId: ticker.getTypeId()},
		worldX:   startX,
		worldY:   startY,
		children: []*Entity{},
		dead:     false,
		ticker:   ticker,
	}
	return &res
}

// An EntityTicker with id 0 that just moves 1 tile to the right each tick. Basically just for testing.
type sillyBeetle struct{}

var BEETLE = &sillyBeetle{}

func (s *sillyBeetle) getTypeId() int32 {
	return 0
}

func (s *sillyBeetle) tick(e *Entity, t *tile) {
	e.worldX += 1
}



