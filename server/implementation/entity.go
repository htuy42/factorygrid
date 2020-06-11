package implementation

import proto "server/protos"

// An entity within the grid
type entity struct{
	// The wire representation of the entity. The client doesn't need state information about the entity, just its
	// type and location so it knows how to draw it.
	internal *proto.Entity
	// The location of the entity in world coordinates.
	worldX, worldY int32
	// During a given tick, all the new entities this one has created. At the end of a tick, this list will be
	// emptied and all the entities in it added to the world.
	children []*entity
	// Whether or not the entity is dead. If this is set to true, the entity will be removed.
	dead bool
	// The ticker that defines specifically what type of entity this is.
	ticker entityTicker
}

// Provides the actual identity of a type of entity and its control logic.
type entityTicker interface{
	// Get the entities type id
	getTypeId() int32
	// Update the entity for a new world tick. Can read the tile, but should generally not modify it. There aren't
	// any specific controls in place to enforce this, so its easy to break things if you want to.
	// todo put entity in a different package and obscure tile private members behind accessor functions so they can't
	// be messed up as easily
	tick(e *entity, t *tile)
}

// Just call the ticker, doesn't need to do anything else.
func (e *entity) tick(t *tile){
	e.ticker.tick(e,t)
}

// Create an entity at the given location with the given ticker
func makeEntity(startX, startY int32, ticker entityTicker) *entity{
	res := entity{
		internal: &proto.Entity{TypeId: ticker.getTypeId()},
		worldX:   startX,
		worldY:   startY,
		children: []*entity{},
		dead:     false,
		ticker:   ticker,
	}
	return &res
}

// An entityTicker with id 0 that just moves 1 tile to the right each tick. Basically just for testing.
type sillyBeetle struct{}

var BEETLE = &sillyBeetle{}

func (s *sillyBeetle) getTypeId() int32 {
	return 0
}

func (s *sillyBeetle) tick(e *entity, t *tile) {
	e.worldX += 1
}



