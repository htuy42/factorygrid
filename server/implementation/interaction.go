package implementation

// An Interaction is a pre-packaged way to respond to a given user input. 
type Interaction interface {
	// The function to run every time the Interaction is triggered.
	trigger(t *tile, c *cell, xLocal, yLocal, xGlobal, yGlobal int32) bool
	// Make a new copy of the Interaction. In practice, when an Interaction is passed to factory, a copy of it will 
	// be made for each tile the factory manages. If an Interaction doesn't care about having a separate instance per 
	// tile and doesn't mind being called multiple times at once (since a single tiles tick is single-threaded but tiles
	// tick in parallel with eachother), it can just return itself from this method. Otherwise, it needs to return 
	// actual copies. Technically it can return other things that are not copies of itself, but this is probably not 
	// usually a good idea.
	makeCopy() Interaction
}

// An interaction that changes the typeId of the tile it is applied to.
type PaintInteraction struct{
	targetTypeId int32
}

func NewPaintInteraction(targetTypeId int32) *PaintInteraction {
	return &PaintInteraction{targetTypeId: targetTypeId}
}

func (p *PaintInteraction) trigger(t *tile, c *cell, xLocal, yLocal, xGlobal, yGlobal int32) bool {
	c.typeId = p.targetTypeId
	return true
}

// This interaction doesn't have any mutable state and doesn't care about parallel calls so it just returns itself on
// copy
func (p *PaintInteraction) makeCopy() Interaction {
	return p
}

type SpawnInteraction struct{
	spawner func(xGlobal,yGlobal int32) EntityTicker
}

func NewSpawnInteraction(spawner func(x, y int32) EntityTicker) *SpawnInteraction {
	return &SpawnInteraction{spawner: spawner}
}

func (s *SpawnInteraction) trigger(t *tile, c *cell, xLocal, yLocal, xGlobal, yGlobal int32) bool {
	newEnt := makeEntity(xGlobal,yGlobal,s.spawner(xGlobal,yGlobal))
	t.addEntity(newEnt)
	return true
}

func (s *SpawnInteraction) makeCopy() Interaction {
	return s
}
