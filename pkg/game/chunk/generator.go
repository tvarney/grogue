package chunk

import (
	"log"
	"math/bits"

	perlin "github.com/aquilax/go-perlin"
	"github.com/tvarney/grogue/pkg/game/material"
	"github.com/tvarney/grogue/pkg/game/simplehash"
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

	surface *perlin.Perlin
}

// NewGenerator creates a new Generator instance with the given materials.
func NewGenerator(seed int64, mats []*material.Material) *Generator {
	log.Printf("game.chunk::NewGenerator(): given %d materials", len(mats))
	g := &Generator{
		Materials: mats,
		Air:       0,
		Water:     1,
		Bedrock:   2,

		surface: perlin.NewPerlin(2, 2, 3, seed),
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

// Flat creates a new flat chunk using the settings from the generator.
func (g *Generator) Flat(cx, cy int64) *Chunk {
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

	hash := simplehash.Initial32.AddInt64(cx).AddInt64(cy)

	idx := uint16(0)
	for n := 0; n < layertiles; n++ {
		c.Tiles[idx] = bedrock
		c.Tiles[idx].Random = uint32(hash.AddUint16(idx))
		idx++
	}

	const stonecount = layertiles * 30
	for n := 0; n < stonecount; n++ {
		c.Tiles[idx] = stone
		c.Tiles[idx].Random = uint32(hash.AddUint16(idx))
		idx++
	}

	const dirtcount = layertiles * 2
	for n := 0; n < dirtcount; n++ {
		c.Tiles[idx] = dirt
		c.Tiles[idx].Random = uint32(hash.AddUint16(idx))
		idx++
	}

	for n := 0; n < layertiles; n++ {
		c.Tiles[idx] = grass
		c.Tiles[idx].Random = uint32(hash.AddUint16(idx))
		idx++
	}

	for ; idx < TileCount; idx++ {
		c.Tiles[idx] = empty
		c.Tiles[idx].Random = uint32(hash.AddUint16(idx))
	}

	return c
}

// Generate creates a new randomized chunk.
func (g *Generator) Generate(cx, cy int64) *Chunk {
	chunk := g.Flat(cx, cy)

	for y := 0; y < Length; y++ {
		for x := 0; x < Width; x++ {
			fx := float64(cx) + float64(x)/Width
			fy := float64(cy) + float64(y)/Length
			n := g.surface.Noise2D(fx, fy)
			switch {
			case n < 0.005:
				// Carve out for water
				// Remove surface floor, clear flags
				t := chunk.Get(x, y, 33)
				t.Floor = tile.Part{
					Definition: tile.FloorEmpty,
					Material:   g.Air,
				}
				t.Flags = 0

				// Remove block from below, add liquid
				t = chunk.Get(x, y, 32)
				t.Block = tile.Part{
					Definition: tile.BlockEmpty,
					Material:   g.Air,
				}
				t.Liquid = 7
				t.LiquidMat = g.Water
			case n < 0.05:
				// Remove grass
				chunk.Get(x, y, 33).Flags &= tile.StateFlags(bits.Reverse16(uint16(tile.HasGrass)))
			}
		}
	}

	return chunk
}
