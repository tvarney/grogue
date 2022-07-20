package material

import "github.com/tvarney/grogue/pkg/game/color"

// ID is a material reference ID.
//
// This is functionally an index to the material in the list of materials.
type ID uint16

// Type is an enumeration of material categories.
//
// This determines which recipies a material may be used in, and how the
// material is used during world generation.
type Type uint16

const (
	Stone Type = iota
	Metal
	Soil
	Sand
	Glass
	Gem
	Wood
	Bone
	Flesh
	Misc
)

// State is a set of values for a material which depend on the physical state
// the material is in.
//
// Valid states are 'liquid', 'solid', and 'gas'. The game doesn't bother with
// plasmas or other more exotic states of matter.
type State struct {
	Name      string
	Adjective string
	Color     color.Enum
}

// Material is the definition of a material for the game.
type Material struct {
	Type   Type
	Solid  State
	Liquid State
	Gas    State
}
