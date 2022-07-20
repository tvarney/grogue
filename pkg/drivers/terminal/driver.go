package terminal

import (
	"io"
	"io/ioutil"
	"log"
	"os"

	"github.com/gdamore/tcell"
	"github.com/tvarney/grogue/pkg/game"
)

// Driver is the terminal driver struct.
type Driver struct {
	blocks []Displayer
	floors []Displayer
	grass  Displayer
	player Displayer

	logfile string
	logfp   io.WriteCloser
	screen  tcell.Screen
	width   int
	height  int
}

// New creates a new Driver with a new game instance.
func New() *Driver {
	return &Driver{
		blocks: DefaultBlocks(),
		floors: DefaultFloors(),
		grass:  Random([]rune{'.', ',', ';'}),
		player: Simple('☺'),
	}
}

func (d *Driver) openlog(filename string) error {
	// Close any previously opened log
	if d.logfp != nil {
		d.logfp.Close()
		d.logfp = nil
	}

	// Update our internal reference to the filename
	d.logfile = filename

	// Now, if the requested filename is empty, set up the discard writer
	if filename == "" {
		log.SetOutput(ioutil.Discard)
		return nil
	}

	// Otherwise, attempt to open the log file
	fp, err := os.Create(filename)
	if err != nil {
		log.SetOutput(ioutil.Discard)
		return err
	}

	log.SetOutput(fp)
	d.logfp = fp
	return nil
}

// SetLog sets the logging behavior of the driver.
//
// As the driver makes use of the standard log package, this will set the
// filename to redirect the log package to. If filename is the empty string,
// this will result in a NopCloser
func (d *Driver) SetLog(filename string) error {
	if d.screen != nil {
		return d.openlog(filename)
	}
	d.logfile = filename
	return nil
}

// Init initializes the terminal driver.
func (d *Driver) Init() error {
	if d.screen != nil {
		return nil
	}

	s, err := tcell.NewScreen()
	if err != nil {
		return err
	}
	if err := s.Init(); err != nil {
		return err
	}

	// Just discard the error; ideally we would report this, but a failure to
	// open the log file shouldn't break our init function either. That makes
	// it a good candidate to log...but we just failed to open the log, so it
	// doesn't matter.
	_ = d.openlog(d.logfile)

	d.screen = s
	d.width, d.height = s.Size()
	return nil
}

// Finalize releases terminal driver resources.
func (d *Driver) Finalize() {
	if d.screen == nil {
		return
	}
	d.screen.Fini()
	d.screen = nil
}

// PollAction gets the next action to update the game with.
func (d *Driver) PollAction(app *game.Application) game.Action {
	var action game.Action
	for action == game.ActionNone {
		ev := d.screen.PollEvent()
		if ev == nil {
			// ev may only be nil if the terminal is not initialized, so
			// handle this as if the 'window' was closed.
			return game.ActionQuit
		}

		switch e := ev.(type) {
		case *tcell.EventResize:
			d.width, d.height = e.Size()
			d.screen.Sync()
			d.Draw(app)
		case *tcell.EventKey:
			menu := app.GetMenu()
			if menu != nil {
				action = d.HandleKeyEventMenu(app, e)
			} else {
				action = d.HandleKeyEventGame(app, e)
			}
		default:
			app.Quit()
			log.Printf("Unknown event <%T>: %v", ev, ev)
		}
	}

	return action
}

func (d *Driver) HandleKeyEventGame(app *game.Application, event *tcell.EventKey) game.Action {
	switch event.Key() {
	case tcell.KeyRune:
		switch event.Rune() {
		case '>':
			return game.ActionMoveUp
		case '<':
			return game.ActionMoveDown
		case 'j':
			return game.ActionMoveSouth
		case 'k':
			return game.ActionMoveNorth
		case 'h':
			return game.ActionMoveWest
		case 'l':
			return game.ActionMoveEast
		case '.':
			return game.ActionWait
		}
	case tcell.KeyLeft:
		return game.ActionMoveWest
	case tcell.KeyRight:
		return game.ActionMoveEast
	case tcell.KeyUp:
		return game.ActionMoveNorth
	case tcell.KeyDown:
		return game.ActionMoveSouth
	case tcell.KeyCtrlC:
		return game.ActionQuit
	}
	return game.ActionNone
}

func (d *Driver) HandleKeyEventMenu(app *game.Application, event *tcell.EventKey) game.Action {
	switch event.Key() {
	case tcell.KeyRune:
		switch event.Rune() {
		case 'j':
			return game.ActionMenuDown
		case 'k':
			return game.ActionMenuUp
		}
	case tcell.KeyEscape:
		return game.ActionMenuClose
	case tcell.KeyDown:
		return game.ActionMenuDown
	case tcell.KeyUp:
		return game.ActionMenuUp
	case tcell.KeyEnter:
		return game.ActionMenuSelect
	case tcell.KeyCtrlC:
		return game.ActionQuit
	}
	return game.ActionNone
}
