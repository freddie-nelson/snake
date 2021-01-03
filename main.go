package main

import (
	tl "github.com/JoelOtter/termloop"
)

// EndLevel the level that is displayed when the game is over
var EndLevel *tl.BaseLevel

// GameOver is the game over
var GameOver = false

// TheFood the food
var TheFood *Food

func main() {
	game := tl.NewGame()

	// setup levels
	mainLevel := SetupMainLevel()
	EndLevel = SetupEndLevel()

	screen := game.Screen()
	screen.SetLevel(mainLevel)
	screen.SetFps(10)

	game.Start()
}
