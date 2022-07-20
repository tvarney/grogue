package game

// Action is an enumeration of actions which the Application handles.
//
// These values are expected to be translated from keyboard and mouse events
// by the driver as appropriate.
type Action int

const (
	ActionQuit Action = iota - 1
	ActionNone
	ActionMoveNorth
	ActionMoveSouth
	ActionMoveEast
	ActionMoveWest
	ActionMoveUp
	ActionMoveDown
	ActionWait
	ActionMenuOpen

	ActionMenuUp
	ActionMenuDown
	ActionMenuLeft
	ActionMenuRight
	ActionMenuSelect
	ActionMenuClose
)
