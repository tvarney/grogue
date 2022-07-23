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

// Get the tile at the given (x,y,z) coordinates.
//
// This function does no bounds checking, and as such may result in a panic if
// the coordinates fall outside of the chunk boundaries. Tiles are stored by
// value in the chunk, while this function returns a pointer to the tile. This
// allows mutation of the tile by the caller of this function. As such, there
// is no corresponding `Set` function.
func (c *Chunk) Get(x, y, z int) *tile.State {
	return &(c.Tiles[(z*LayerSize)+(y*Width)+x])
}
