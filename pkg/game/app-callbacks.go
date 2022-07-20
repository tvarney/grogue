package game

import "log"

// AddMenu adds a menu to the registered list of menus.
func (a *Application) AddMenu(m Menu) {
	a.menus[m.GetID()] = m
}

// PushMenu pushes the menu with the given ID onto the menu stack.
func (a *Application) PushMenu(id string) {
	m, ok := a.menus[id]
	if ok {
		log.Printf("game.Application::PushMenu(%q): Pushing menu onto stack", id)
		if len(a.menu) > 0 {
			a.menu[len(a.menu)-1].Pause(a)
		}
		a.menu = append(a.menu, m)
		m.Start(a)
	} else {
		log.Printf("game.Application::PushMenu(%q): No menu with the given ID", id)
	}
}

// PopMenu removes the topmost menu from the menu stack.
func (a *Application) PopMenu() {
	if len(a.menu) == 0 {
		return
	}

	a.menu[len(a.menu)-1].Stop(a)
	a.menu = a.menu[:len(a.menu)-1]
	if len(a.menu) > 0 {
		a.menu[len(a.menu)-1].Resume(a)
	}
}

// GetMenu returns the current menu.
func (a *Application) GetMenu() Menu {
	if len(a.menu) > 0 {
		return a.menu[len(a.menu)-1]
	}
	return nil
}

// Quit signals that the Application should be shut down.
func (a *Application) Quit() {
	log.Printf("game.Application::Quit()")
	a.Running = false
}
