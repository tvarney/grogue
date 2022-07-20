package game

// RenderRequest is an enumeration of render requests for the game driver to
// handle.
type RenderRequest uint32

const (
	RenderNoChange RenderRequest = iota
	RenderIncremental
	RenderFull
)
