package main

import tl "github.com/JoelOtter/termloop"

// Listener a listener to execute game logic each frame
type Listener struct {
	*tl.Entity
}

// NewListener creates a new listener
func NewListener() *Listener {
	return &Listener{tl.NewEntity(0, 0, 0, 0)}
}

// Draw executes every frame to perform game logic based on the games current state
func (listener *Listener) Draw(screen *tl.Screen) {
	if GameOver && screen.Level() != EndLevel {
		text := tl.NewEntityFromCanvas(ScreenWidth/2-39, ScreenHeight/2-4, tl.CanvasFromString(GameOverText))
		EndLevel.AddEntity(text)
		screen.SetLevel(EndLevel)
	}
}
