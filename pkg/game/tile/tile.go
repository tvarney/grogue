package tile

import "github.com/tvarney/grogue/pkg/game/material"

const (
	HasGrass StateFlags = 1 << iota
)

// ID is the tile-definition ID for a tile-state.
type ID uint16

// Definition is a tile definition.
//
// This type holds the static information about what a tile is.
type Definition struct {
	Name string
}

// StateFlags is a bitfield of flags for a tile state.
type StateFlags uint16

// Part defines a material part of a tile.
type Part struct {
	Definition ID
	Material   material.ID
}

// State represents a concrete tile.
type State struct {
	Block  Part
	Floor  Part
	Flags  StateFlags
	Value  uint16
	Random uint32
}
