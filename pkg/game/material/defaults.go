package material

import (
	"github.com/tvarney/grogue/pkg/game/color"
)

// DefaultMaterials returns an array of materials.
//
// This is a placeholder function; eventually this will be replaced by reading
// data files, with only the Air, Bedrock, and Water materials being
// pre-defined.
func DefaultMaterials() []*Material {
	return []*Material{
		{
			Type: Misc,
			Solid: State{
				Name:      "solid air",
				Adjective: "solid air",
				Color:     color.BrightWhite,
			},
			Liquid: State{
				Name:      "liquid air",
				Adjective: "liquid air",
				Color:     color.BrightWhite,
			},
			Gas: State{
				Name:      "air",
				Adjective: "air",
				Color:     color.BrightWhite,
			},
		},
		{
			Type: Stone,
			Solid: State{
				Name:      "ice",
				Adjective: "ice",
				Color:     color.BrightCyan,
			},
			Liquid: State{
				Name:      "water",
				Adjective: "water",
				Color:     color.Blue,
			},
			Gas: State{
				Name:      "steam",
				Adjective: "steam",
				Color:     color.White,
			},
		},
		{
			Type: Stone,
			Solid: State{
				Name:      "bedrock",
				Adjective: "bedrock",
				Color:     color.Black,
			},
			Liquid: State{
				Name:      "magma",
				Adjective: "magma",
				Color:     color.BrightOrange,
			},
			Gas: State{
				Name:      "vaporized bedrock",
				Adjective: "vaporized bedrock",
				Color:     color.BrightOrange,
			},
		},
		{
			Type: Stone,
			Solid: State{
				Name:      "stone",
				Adjective: "stone",
				Color:     color.Gray,
			},
			Liquid: State{
				Name:      "lava",
				Adjective: "lava",
				Color:     color.Orange,
			},
			Gas: State{
				Name:      "vaporized stone",
				Adjective: "vaporized stone",
				Color:     color.BrightOrange,
			},
		},
		{
			Type: Soil,
			Solid: State{
				Name:      "dirt",
				Adjective: "dirt",
				Color:     color.Brown,
			},
			Liquid: State{
				Name:      "dirt",
				Adjective: "dirt",
				Color:     color.Brown,
			},
			Gas: State{
				Name:      "dirt",
				Adjective: "dirt",
				Color:     color.Brown,
			},
		},
	}
}
