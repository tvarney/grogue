package game

type ChunkCoords struct {
	X int
	Y int
}

type Coords struct {
	X int
	Y int
	Z int

	Chunk ChunkCoords
}
