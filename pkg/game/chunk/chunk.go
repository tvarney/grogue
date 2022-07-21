package chunk

import (
	"github.com/tvarney/grogue/pkg/game/tile"
)

const (
	// Width is the width (x) of each chunk
	Width = 32

	// Height is the height (z) of each chunk
	Height = 62

	// Length is the length (y) of each chunk
	Length = 32

	LayerSize = Width * Length

	// TileCount is the total number of tile.State entries in a chunk.
	//
	// This is the product of the chunk Width, Height, and Length. This value
	// should be kept below the maximum value for a uint16 index.
	//
	// This is kept around for rendering code; for some reason it doesn't want
	// to work for the actual Chunk definition.
	TileCount = Width * Height * Length
)

// Chunk is a chunk of the game world.
type Chunk struct {
	Tiles [Width * Height * Length]tile.State
}

// New returns a new Chunk instance.
func New() *Chunk {
	return &Chunk{}
}

func (c *Chunk) Get(x, y, z int) *tile.State {
	return &(c.Tiles[(z*LayerSize)+(y*Width)+x])
}
