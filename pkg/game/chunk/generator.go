package chunk

import (
	"log"

	"github.com/tvarney/grogue/pkg/game/material"
	"github.com/tvarney/grogue/pkg/game/tile"
)

// Generator builds new chunks for the world as needed.
type Generator struct {
	Materials []*material.Material

	Air     material.ID
	Bedrock material.ID
	Water   material.ID
	Stone   []material.ID
	Soil    []material.ID
}

// NewGenerator creates a new Generator instance with the given materials.
func NewGenerator(mats []*material.Material) *Generator {
	log.Printf("game.chunk::NewGenerator(): given %d materials", len(mats))
	g := &Generator{
		Materials: mats,
		Air:       0,
		Water:     1,
		Bedrock:   2,
	}

	// Skip the first 3 (which must always exist), but otherwise categorize
	// our materials.
	for i, m := range mats[3:] {
		switch m.Type {
		case material.Stone:
			log.Printf("Using material.ID(%d) as a stone", i+3)
			g.Stone = append(g.Stone, material.ID(i+3))
		case material.Soil:
			log.Printf("Using material.ID(%d) as a soil", i+3)
			g.Soil = append(g.Soil, material.ID(i+3))
		}
	}

	return g
}

// Generate creates a new chunk using the settings from the generator.
//
// Right now this just generates a simple flat chunk.
func (g *Generator) Generate() *Chunk {
	const layertiles = Width * Length

	bedrock := tile.State{
		Block: tile.Part{Definition: tile.BlockStone, Material: g.Bedrock},
		Floor: tile.Part{Definition: tile.FloorStone, Material: g.Bedrock},
	}
	log.Printf("game.chunk.Generator::Generate(): Bedrock %#v", bedrock)
	stone := tile.State{
		Block: tile.Part{Definition: tile.BlockStone, Material: g.Stone[0]},
		Floor: tile.Part{Definition: tile.FloorStone, Material: g.Stone[0]},
	}
	log.Printf("game.chunk.Generator::Generate(): Stone %#v", stone)
	dirt := tile.State{
		Block: tile.Part{Definition: tile.BlockSoil, Material: g.Soil[0]},
		Floor: tile.Part{Definition: tile.FloorSoil, Material: g.Soil[0]},
	}
	log.Printf("game.chunk.Generator::Generate(): Dirt %#v", dirt)
	grass := tile.State{
		Block: tile.Part{Definition: tile.BlockEmpty, Material: g.Air},
		Floor: tile.Part{Definition: tile.FloorSoil, Material: g.Soil[0]},
		Flags: tile.HasGrass,
	}
	log.Printf("game.chunk.Generator::Generate(): Grass %#v", grass)
	empty := tile.State{
		Block: tile.Part{Definition: tile.BlockEmpty, Material: g.Air},
		Floor: tile.Part{Definition: tile.FloorEmpty, Material: g.Air},
	}
	log.Printf("game.chunk.Generator::Generate(): Empty %#v", empty)

	c := &Chunk{}

	idx := 0
	for n := 0; n < layertiles; n++ {
		c.Tiles[idx] = bedrock
		idx++
	}

	const stonecount = layertiles * 30
	for n := 0; n < stonecount; n++ {
		c.Tiles[idx] = stone
		idx++
	}

	const dirtcount = layertiles * 2
	for n := 0; n < dirtcount; n++ {
		c.Tiles[idx] = dirt
		idx++
	}

	for n := 0; n < layertiles; n++ {
		c.Tiles[idx] = grass
		idx++
	}

	for ; idx < TileCount; idx++ {
		c.Tiles[idx] = empty
	}

	return c
}