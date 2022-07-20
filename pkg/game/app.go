package game

import (
	"log"

	"github.com/tvarney/grogue/pkg/game/chunk"
	"github.com/tvarney/grogue/pkg/game/material"
)

// Application is the main game struct, which holds application and game
// states.
type Application struct {
	Running   bool
	InGame    bool
	Materials []*material.Material
	Chunk     *chunk.Chunk
	Generator *chunk.Generator

	PlayerX int
	PlayerY int
	PlayerZ int

	menu  []Menu
	menus map[string]Menu
}

// New returns a new Application instance.
func New() *Application {
	mats := material.DefaultMaterials()
	log.Printf("game::New(): Using %d materials", len(mats))
	app := &Application{
		Running:   true,
		InGame:    false,
		Materials: mats,
		Chunk:     nil,
		Generator: chunk.NewGenerator(mats),
		PlayerZ:   0,

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
	// Unconditionally handle the quit signal
	if action == ActionQuit {
		log.Printf("game.Application::Update(): Handling ActionQuit")
		a.Quit()
		return RenderNoChange
	}

	// If we have a menu, just let it handle the action
	if len(a.menu) > 0 {
		log.Printf("game.Application::Update(): Handling action via menu")
		return a.menu[len(a.menu)-1].HandleAction(action, a)
	}

	// If not in the game and no menus, don't do anything and signal we should
	// exit
	if !a.InGame {
		log.Printf("game.Application::Update(): No menu and not in game; quitting")
		a.Quit()
		return RenderNoChange
	}

	// Handle game actions
	switch action {
	case ActionMoveDown:
		if a.PlayerZ > 0 {
			a.PlayerZ--
			return RenderIncremental
		}
	case ActionMoveUp:
		if a.PlayerZ < chunk.Height-1 {
			a.PlayerZ++
			return RenderIncremental
		}
	case ActionMoveNorth:
		if a.PlayerY > 0 {
			a.PlayerY--
			return RenderIncremental
		}
	case ActionMoveSouth:
		if a.PlayerY < chunk.Length-1 {
			a.PlayerY++
			return RenderIncremental
		}
	case ActionMoveWest:
		if a.PlayerX > 0 {
			a.PlayerX--
			return RenderIncremental
		}
	case ActionMoveEast:
		if a.PlayerX < chunk.Width-1 {
			a.PlayerX++
			return RenderIncremental
		}
	}
	return RenderNoChange
}
