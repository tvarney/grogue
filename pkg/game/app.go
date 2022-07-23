package game

import (
	"log"

	"github.com/tvarney/grogue/pkg/game/chunk"
	"github.com/tvarney/grogue/pkg/game/material"
	"github.com/tvarney/grogue/pkg/game/tile"
)

// Application is the main game struct, which holds application and game
// states.
type Application struct {
	Running      bool
	InGame       bool
	Materials    []*material.Material
	Blocks       []tile.Definition
	Floors       []tile.Definition
	ActiveChunks [9]*chunk.Chunk
	Generator    *chunk.Generator

	Player Coords

	menu  []Menu
	menus map[string]Menu
}

// New returns a new Application instance.
func New() *Application {
	mats := material.DefaultMaterials()
	log.Printf("game::New(): Using %d materials", len(mats))
	blocks, floors := tile.DefaultDefinitions()
	app := &Application{
		Running:   true,
		InGame:    false,
		Materials: mats,
		Blocks:    blocks,
		Floors:    floors,
		Generator: chunk.NewGenerator(mats),

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
		return a.UpdateMovePlayer(0, 0, -1)
	case ActionMoveUp:
		return a.UpdateMovePlayer(0, 0, 1)
	case ActionMoveNorth:
		return a.UpdateMovePlayer(0, -1, 0)
	case ActionMoveSouth:
		return a.UpdateMovePlayer(0, 1, 0)
	case ActionMoveWest:
		return a.UpdateMovePlayer(-1, 0, 0)
	case ActionMoveEast:
		return a.UpdateMovePlayer(1, 0, 0)
	}
	return RenderNoChange
}
