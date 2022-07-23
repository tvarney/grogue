package game

import "github.com/tvarney/grogue/pkg/game/chunk"

// UpdateMovePlayer implements player movement.
//
// This function takes a delta-x, delta-y, and delta-z value; these values
// are assumed to be one of -1, 0, or 1. All other values are not handled
// (yet).
func (a *Application) UpdateMovePlayer(dx, dy, dz int) RenderRequest {
	ret := RenderNoChange

	switch dx {
	case -1:
		if a.Player.X == 0 {
			if a.Player.Chunk.X != -1 {
				a.Player.X = chunk.Width - 1
				a.Player.Chunk.X--
				ret = RenderIncremental
			}
		} else {
			a.Player.X--
			ret = RenderIncremental
		}
	case 1:
		if a.Player.X == chunk.Width-1 {
			if a.Player.Chunk.X != 1 {
				a.Player.X = 0
				a.Player.Chunk.X++
				ret = RenderIncremental
			}
		} else {
			a.Player.X++
			ret = RenderIncremental
		}
	}

	switch dy {
	case -1:
		if a.Player.Y == 0 {
			if a.Player.Chunk.Y != -1 {
				a.Player.Y = chunk.Length - 1
				a.Player.Chunk.Y--
				ret = RenderIncremental
			}
		} else {
			a.Player.Y--
			ret = RenderIncremental
		}
	case 1:
		if a.Player.Y == chunk.Width-1 {
			if a.Player.Chunk.Y != 1 {
				a.Player.Y = 0
				a.Player.Chunk.Y++
				ret = RenderIncremental
			}
		} else {
			a.Player.Y++
			ret = RenderIncremental
		}
	}

	switch dz {
	case -1:
		if a.Player.Z > 0 {
			a.Player.Z--
			ret = RenderIncremental
		}
	case 1:
		if a.Player.Z < chunk.Height-1 {
			a.Player.Z++
			ret = RenderIncremental
		}
	}

	return ret
}
