package game

import "log"

const (
	MainMenuID = "main-menu"
)

// Menu defines how game drivers may interact with menus in the game.
type Menu interface {
	Start(*Application)
	Stop(*Application)
	Pause(*Application)
	Resume(*Application)
	GetID() string
	GetTitle() string
	GetOptions() []string
	GetOption() int
	SetOption(int)
	HandleAction(Action, *Application) RenderRequest
}

// StaticMenu is a menu which is fully defined statically.
type StaticMenu struct {
	ID       string
	Title    string
	OnStart  func(*Application)
	OnStop   func(*Application)
	OnPause  func(*Application)
	OnResume func(*Application)
	Options  []string
	Actions  []func(*Application) RenderRequest
	Cursor   int

	instances int
}

// Start initializes the menu for display.
//
// This generally means that the cursor will be set to the initial position.
func (s *StaticMenu) Start(app *Application) {
	s.instances++
	if s.instances == 1 {
		s.Cursor = 0
		if s.OnStart != nil {
			log.Printf("game.StaticMenu::Start(): ID: %q", s.ID)
			s.OnStart(app)
		}
	} else {
		if s.OnResume != nil {
			log.Printf("game.StaticMenu::Resume(): ID: %q", s.ID)
			s.OnResume(app)
		}
	}
}

// Stop finalizes the menu.
func (s *StaticMenu) Stop(app *Application) {
	s.instances--
	if s.instances == 0 {
		log.Printf("game.StaticMenu::Stop(): ID: %q", s.ID)
		if s.OnStop != nil {
			s.OnStop(app)
		}
	} else {
		log.Printf("game.StaticMenu::Pause(): ID: %q", s.ID)
		if s.OnPause != nil {
			s.OnPause(app)
		}
	}
}

// Pause pauses the menu.
//
// This is called when a menu is pushed 'on top of' this one.
func (s *StaticMenu) Pause(app *Application) {
	log.Printf("game.StaticMenu::Pause(): ID: %q", s.ID)
	if s.OnPause != nil {
		s.OnPause(app)
	}
}

// Resume resumes the menu.
//
// This is called when a menu is 'revealed' by a push action.
func (s *StaticMenu) Resume(app *Application) {
	log.Printf("game.StaticMenu::Resume(): ID: %q", s.ID)
	if s.OnResume != nil {
		s.OnResume(app)
	}
}

// GetID returns the ID of the menu.
func (s *StaticMenu) GetID() string {
	return s.ID
}

// GetTitle returns the title of the menu.
func (s *StaticMenu) GetTitle() string {
	return s.Title
}

// GetOptions returns the options the menu defines.
func (s *StaticMenu) GetOptions() []string {
	return s.Options
}

// GetOption returns the currently selected option.
func (s *StaticMenu) GetOption() int {
	return s.Cursor
}

// SetOption sets the currently focused option.
func (s *StaticMenu) SetOption(idx int) {
	if idx < 0 {
		idx = 0
	} else if idx >= len(s.Options) {
		idx = len(s.Options) - 1
	}
	s.Cursor = idx
}

// HandleAction handles menu actions.
func (s *StaticMenu) HandleAction(a Action, app *Application) RenderRequest {
	switch a {
	case ActionMenuUp:
		if s.Cursor > 0 {
			s.Cursor--
			return RenderIncremental
		}
	case ActionMenuDown:
		if s.Cursor < len(s.Options)-1 {
			s.Cursor++
			return RenderIncremental
		}
	case ActionMenuSelect:
		if s.Cursor < 0 || s.Cursor >= len(s.Actions) {
			return RenderNoChange
		}
		callback := s.Actions[s.Cursor]
		if callback != nil {
			return callback(app)
		}
	case ActionMenuClose:
		app.PopMenu()
		return RenderFull
	}

	// Fallthrough from one of the above, or an unhandled action
	return RenderNoChange
}

// NewMainMenu returns a new StaticMenu instance for the main menu.
func NewMainMenu() *StaticMenu {
	return &StaticMenu{
		ID:    "main-menu",
		Title: "Main Menu",
		Options: []string{
			"New Game",
			"Load Game",
			"Options",
			"Packages",
			"Quit",
		},
		Actions: []func(*Application) RenderRequest{
			func(app *Application) RenderRequest {
				log.Printf("game.StaticMenu::Actions[0]: Selected New Game")
				app.InGame = true
				app.Chunk = app.Generator.Generate()
				app.PopMenu()
				return RenderFull
			},
			nil,
			nil,
			nil,
			func(app *Application) RenderRequest {
				app.Quit()
				return RenderFull
			},
		},
	}
}
