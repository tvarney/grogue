package terminal

import (
	"github.com/tvarney/grogue/pkg/game/tile"
)

// Displayer is the interface all tile displayers must implement.
type Displayer interface {
	Rune(cx, cy int, x, y uint16, t *tile.State) rune
}

// Simple is a displayer which only ever returns a single rune.
type Simple rune

func (s Simple) Rune(cx, cy int, x, y uint16, t *tile.State) rune {
	return rune(s)
}

// Random is a displayer which chooses a 'pseudo-random' tile to display based
// on the coordinates.
type Random []rune

func (r Random) Rune(cx, cy int, x, y uint16, t *tile.State) rune {
	return r[int(t.Random)%len(r)]
}

// LiquidNumber is a displayer for liquids which shows their depth.
type LiquidNumber struct{}

func (l LiquidNumber) Rune(cx, cy int, x, y uint16, t *tile.State) rune {
	return rune(uint32('0') + uint32(t.Liquid&0x0007))
}

func DefaultBlocks() []Displayer {
	return []Displayer{
		Simple(' '),
		Simple('█'),
		Simple('▓'),
		Simple('█'),
		Simple('█'),
	}
}

func DefaultFloors() []Displayer {
	return []Displayer{
		Simple(' '),
		Simple('.'),
		Simple('.'),
		Simple('.'),
		Simple('.'),
	}
}
