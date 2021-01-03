package main

import tl "github.com/JoelOtter/termloop"

// SetupMainLevel creates the main level, adds all entities for the main level to it
// and returns a pointer to the main level
func SetupMainLevel() *tl.BaseLevel {
	mainLevel := tl.NewBaseLevel(tl.Cell{
		Bg: tl.ColorBlack,
		Ch: ' ',
	})

	snake := NewSnake()
	mainLevel.AddEntity(snake)

	TheFood = NewFood()
	mainLevel.AddEntity(TheFood)

	scoreText := NewScoreText(tl.ColorCyan, tl.ColorBlack)
	mainLevel.AddEntity(scoreText)

	listener := NewListener()
	mainLevel.AddEntity(listener)

	return mainLevel
}
