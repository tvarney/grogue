package terminal

import (
	"fmt"

	"github.com/gdamore/tcell"
	"github.com/tvarney/grogue/pkg/game"
	"github.com/tvarney/grogue/pkg/game/chunk"
	"github.com/tvarney/grogue/pkg/game/color"
	"github.com/tvarney/grogue/pkg/game/tile"
)

var (
	titleStyle    = tcell.StyleDefault.Bold(true).Underline(true)
	optionStyle   = tcell.StyleDefault
	selectedStyle = tcell.StyleDefault.Underline(true).Foreground(tcell.ColorTeal)

	grassStyles = []tcell.Style{
		tcell.StyleDefault.Foreground(tcell.NewHexColor(int32(color.DarkGreen.Value()))),
		tcell.StyleDefault.Foreground(tcell.NewHexColor(int32(color.Green.Value()))),
	}
	emptyStyle = tcell.StyleDefault.
			Foreground(tcell.NewHexColor(0x001111)).
			Background(tcell.NewHexColor(0x002222))
)

const (
	cursor      = "* "
	eraseCursor = "  "
)

func (d *Driver) Clear() {
	d.screen.Clear()
}

func (d *Driver) Draw(app *game.Application) {
	menu := app.GetMenu()
	if menu != nil {
		d.drawMenu(app, menu)
		return
	}

	if !app.InGame {
		app.Quit()
		return
	}
	d.drawGame(app)
}

func (d *Driver) drawGame(app *game.Application) {
	const layersize = chunk.Width * chunk.Length

	for y := uint16(0); y < chunk.Length; y++ {
		if int(y) >= d.height {
			break
		}
		for x := uint16(0); x < chunk.Width; x++ {
			if int(x) >= d.width {
				break
			}

			t := app.Chunk.Tiles[layersize*app.PlayerZ]
			block := t.Block.Definition
			floor := t.Floor.Definition
			if block == tile.BlockEmpty {
				if floor == tile.FloorEmpty {
					d.screen.SetContent(int(x), int(y), '.', nil, emptyStyle)
				} else {
					if t.Flags&tile.HasGrass != 0 {
						gc := grassStyles[(int(x)*3+int(y)*2)%len(grassStyles)]
						d.screen.SetContent(int(x), int(y), d.grass.Rune(0, 0, x, y, t.Value), nil, gc)
					} else {
						mat := app.Materials[t.Floor.Material]
						d.screen.SetContent(
							int(x), int(y),
							d.floors[floor].Rune(0, 0, x, y, t.Value), nil,
							tcell.StyleDefault.Foreground(tcell.NewHexColor(int32(mat.Solid.Color.Value()))),
						)
					}
				}
			} else {
				mat := app.Materials[t.Block.Material]
				d.screen.SetContent(
					int(x), int(y),
					d.blocks[block].Rune(0, 0, x, y, t.Value), nil,
					tcell.StyleDefault.Foreground(tcell.NewHexColor(int32(mat.Solid.Color.Value()))),
				)
			}
		}
		d.drawString(0, chunk.Length, fmt.Sprintf("Z: %d   ", app.PlayerZ), tcell.StyleDefault)

		d.screen.Show()
	}
}

func (d *Driver) drawMenu(app *game.Application, menu game.Menu) {
	d.drawStringCentered(0, menu.GetTitle(), titleStyle)

	opts := menu.GetOptions()
	maxlen := 0
	for _, o := range opts {
		if len(o) > maxlen {
			maxlen = len(o)
		}
	}
	opt_x := (d.width - (maxlen + len(cursor))) / 2

	selected := menu.GetOption()
	for i, option := range menu.GetOptions() {
		if i == selected {
			d.drawString(opt_x, 2+i, cursor+option, selectedStyle)
		} else {
			d.drawString(opt_x, 2+i, eraseCursor+option, optionStyle)
		}
	}

	d.screen.Show()
}

func (d *Driver) drawString(x, y int, str string, style tcell.Style) {
	if y >= d.height {
		return
	}

	// Figure out _where_ in the string to start
	start := 0
	if x < 0 {
		start = -x
	}
	// If that's past the end, nothing to draw
	if start > len(str) {
		return
	}

	// Calculate how much of the string to draw; if x is negative, this will
	// 'draw' more than the screen width by skipping the start bits.
	n := d.width - x
	if len(str) < n {
		n = len(str)
	}

	for i, r := range str[start:n] {
		d.screen.SetContent(i+x, y, r, nil, style)
	}
}

func (d *Driver) drawStringCentered(y int, str string, style tcell.Style) {
	d.drawString((d.width-len(str))/2, y, str, style)
}
