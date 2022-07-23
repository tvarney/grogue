package game

// ChunkCoords is the (x,y) pair of coordinates identifying a chunk.
type ChunkCoords struct {
	X int
	Y int
}

// Coords is the (x,y,z) tuple of coordinates, along with chunk coordinates.
//
// This type uniquely identifies a position within the game. Valid values for
// the (x,y,z) tuple are ([0-chunk.Width], [0-chunk.Length], [0-chunk.Height]).
// Values outside those ranges are not checked.
type Coords struct {
	X int
	Y int
	Z int

	Chunk ChunkCoords
}
