package tile

import (
	"log"
	"strings"
	"text/template"

	"github.com/tvarney/grogue/pkg/game/material"
)

const (
	HasGrass StateFlags = 1 << iota
)

// ID is the tile-definition ID for a tile-state.
type ID uint16

// Definition is a tile definition.
//
// This type holds the static information about what a tile is.
type Definition struct {
	ID   string
	Name string

	nametemplate *template.Template
}

// GetName returns the name of the tile represented by the given state.
func (d *Definition) GetName(mat *material.Material) string {
	if d.nametemplate == nil {
		d.nametemplate = template.New(d.ID)
		_, err := d.nametemplate.Parse(d.Name)
		if err != nil {
			log.Printf("Error parsing name template for %s", d.ID)
		}
	}

	sb := strings.Builder{}
	d.nametemplate.Execute(&sb, mat)
	return sb.String()
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
	Block     Part
	Floor     Part
	Flags     StateFlags
	Value     uint16
	Liquid    uint16
	LiquidMat material.ID
	Random    uint32
}

// Describe gets a descriptive name of the tile.
func (s *State) Describe(blocks, floors []Definition, mats []*material.Material) string {
	if s.Liquid > 0 {
		return mats[s.LiquidMat].Liquid.Name
	}
	if s.Flags&HasGrass != 0 {
		return "grass"
	}

	if s.Block.Definition != BlockEmpty {
		return (&blocks[s.Block.Definition]).GetName(mats[s.Block.Material])
	}
	if s.Floor.Definition != BlockEmpty {
		return (&floors[s.Floor.Definition]).GetName(mats[s.Floor.Material])
	}
	return "empty"
}
