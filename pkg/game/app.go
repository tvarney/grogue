package game

// Application is the main game struct, which holds application and game
// states.
type Application struct {
	Running bool

	menu  []Menu
	menus map[string]Menu
}

// New returns a new Application instance.
func New() *Application {
	app := &Application{
		Running: true,

		menu:  make([]Menu, 0, 10),
		menus: map[string]Menu{},
	}

	app.AddMenu(NewMainMenu())
	app.PushMenu(MainMenuID)

	return app
}

// Update takes an action from the game driver and updates the state to
// reflect the results of that action.
func (a *Application) Update(action Action) RenderRequest {
	if action == ActionQuit || len(a.menu) == 0 {
		a.Quit()
		return RenderNoChange
	}

	return a.menu[len(a.menu)-1].HandleAction(action, a)
}
