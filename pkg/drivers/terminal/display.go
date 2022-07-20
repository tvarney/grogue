package terminal

import (
	"github.com/gdamore/tcell"
	"github.com/tvarney/grogue/pkg/game"
)

var (
	titleStyle    = tcell.StyleDefault.Bold(true).Underline(true)
	optionStyle   = tcell.StyleDefault
	selectedStyle = tcell.StyleDefault.Underline(true).Foreground(tcell.ColorTeal)
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
	if menu == nil {
		app.Quit()
		return
	}

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
