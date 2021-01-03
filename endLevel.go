package main

import tl "github.com/JoelOtter/termloop"

// SetupEndLevel creates the game over screen and returns a pointer to the level
func SetupEndLevel() *tl.BaseLevel {
	l := tl.NewBaseLevel(tl.Cell{
		Bg: tl.ColorRed,
		Fg: tl.ColorWhite,
		Ch: ' ',
	})

	scoreText := NewScoreText(tl.ColorWhite, tl.ColorRed)
	l.AddEntity(scoreText)

	return l
}
