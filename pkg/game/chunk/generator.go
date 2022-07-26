package chunk

import (
	"log"
	"math"
	"math/bits"
	"math/rand"

	perlin "github.com/aquilax/go-perlin"
	"github.com/tvarney/grogue/pkg/game/material"
	"github.com/tvarney/grogue/pkg/game/simplehash"
	"github.com/tvarney/grogue/pkg/game/tile"
)

const CaveOffset = 5

type CaveParams struct {
	Threshold  float64
	Scale      float64
	TranslateX float64
	TranslateY float64
	CosTheta   float64
	SinTheta   float64
}

func NewCaveParams(r *rand.Rand) CaveParams {
	threshold := 0.1 + r.Float64()*0.3 // Maximum is ~35% cave
	scale := 0.5 + r.Float64()
	angle := r.Float64() * math.Pi * 2
	translateX := r.Float64() * 100.0
	translateY := r.Float64() * 100.0
	return CaveParams{
		Threshold:  threshold,
		Scale:      scale,
		TranslateX: translateX,
		TranslateY: translateY,
		CosTheta:   math.Cos(angle),
		SinTheta:   math.Sin(angle),
	}
}

func (c *CaveParams) GetCoords(cx, cy int64, x, y uint16) (float64, float64) {
	fx := (float64(cx) + float64(x)/Width) + c.TranslateX
	fy := (float64(cy) + float64(y)/Length) + c.TranslateY

	fx2 := (fx*c.CosTheta - fy*c.SinTheta) * c.Scale
	fy2 := (fx*c.SinTheta - fy*c.CosTheta) * c.Scale
	return fx2, fy2
}

func (c *CaveParams) IsCave(p *perlin.Perlin, cx, cy int64, x, y uint16) bool {
	fx, fy := c.GetCoords(cx, cy, x, y)
	return (p.Noise2D(fx, fy)+1.0)/2.0 < c.Threshold
}

// Generator builds new chunks for the world as needed.
type Generator struct {
	Materials []*material.Material

	Air     material.ID
	Bedrock material.ID
	Water   material.ID
	Stone   []material.ID
	Soil    []material.ID

	rand       *rand.Rand
	surface    *perlin.Perlin
	caves      *perlin.Perlin
	caveParams [20]CaveParams
}

// NewGenerator creates a new Generator instance with the given materials.
func NewGenerator(seed int64, mats []*material.Material) *Generator {
	log.Printf("game.chunk::NewGenerator(): given %d materials", len(mats))
	r := rand.New(rand.NewSource(seed))
	g := &Generator{
		Materials: mats,
		Air:       0,
		Water:     1,
		Bedrock:   2,

		rand:    r,
		surface: perlin.NewPerlin(2, 2, 3, r.Int63()),
		caves:   perlin.NewPerlin(1.9, 2, 8, r.Int63()),
	}
	for i := 0; i < len(g.caveParams); i++ {
		g.caveParams[i] = NewCaveParams(r)
		log.Printf(
			"CaveParams[%d] = {\n  Threshold: %0.5f\n  ScaleX: %0.5f\n  ScaleY: %0.5f\n  CosTheta: %0.5f\n  SinTheta: %0.5f\n}",
			i+CaveOffset, g.caveParams[i].Threshold, g.caveParams[i].Scale, g.caveParams[i].Scale,
			g.caveParams[i].CosTheta, g.caveParams[i].SinTheta,
		)
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
	stone := tile.State{
		Block: tile.Part{Definition: tile.BlockStone, Material: g.Stone[0]},
		Floor: tile.Part{Definition: tile.FloorStone, Material: g.Stone[0]},
	}
	dirt := tile.State{
		Block: tile.Part{Definition: tile.BlockSoil, Material: g.Soil[0]},
		Floor: tile.Part{Definition: tile.FloorSoil, Material: g.Soil[0]},
	}
	grass := tile.State{
		Block: tile.Part{Definition: tile.BlockEmpty, Material: g.Air},
		Floor: tile.Part{Definition: tile.FloorSoil, Material: g.Soil[0]},
		Flags: tile.HasGrass,
	}
	empty := tile.State{
		Block: tile.Part{Definition: tile.BlockEmpty, Material: g.Air},
		Floor: tile.Part{Definition: tile.FloorEmpty, Material: g.Air},
	}

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
			// Noise, normalized to [0.0, 1.0]. This is a guassian distribution
			n := (g.surface.Noise2D(fx, fy) + 1.0) / 2.0

			switch {
			case n < 0.35:
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
			case n < 0.37:
				// Remove grass
				chunk.Get(x, y, 33).Flags &= tile.StateFlags(bits.Reverse16(uint16(tile.HasGrass)))
			}
		}
	}

	// Carve out caves; we use a single source of noise, but we rotate,
	// translate, and scale by random values based on the Z level.
	for z := 0; z < len(g.caveParams); z++ {
		for y := uint16(0); y < Length; y++ {
			for x := uint16(0); x < Width; x++ {
				if g.caveParams[z].IsCave(g.caves, cx, cy, x, y) {
					t := chunk.Get(int(x), int(y), int(z)+CaveOffset)
					t.Block.Definition = tile.BlockEmpty
					t.Block.Material = g.Air
				}
			}
		}
	}

	return chunk
}
