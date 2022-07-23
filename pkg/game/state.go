package game

import (
	"github.com/tvarney/grogue/pkg/game/chunk"
	"github.com/tvarney/grogue/pkg/game/material"
	"github.com/tvarney/grogue/pkg/game/tile"
)

// RenderRequest is an enumeration of render requests for the game driver to
// handle.
type RenderRequest uint32

const (
	RenderNoChange RenderRequest = iota
	RenderIncremental
	RenderFull
)

type Game struct {
	Materials []*material.Material
	Blocks    []tile.Definition
	Floors    []tile.Definition

	Player       Coords
	ActiveChunks [9]*chunk.Chunk
	Generator    *chunk.Generator
}
