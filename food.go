package main

import (
	"math/rand"

	tl "github.com/JoelOtter/termloop"
)

// Food snake food yum
type Food struct {
	Coord
}

// NewFood create a new food
func NewFood() *Food {
	return &Food{Coord{10, 1}}
}

// Eat add to score and generate new food
func (food *Food) Eat() {
	Score++

	food.x = randInt(ScreenWidth)
	food.y = randInt(ScreenHeight)
}

// Draw renders the food on screen
func (food *Food) Draw(screen *tl.Screen) {
	screen.RenderCell(food.x, food.y, &tl.Cell{Fg: tl.ColorRed, Bg: tl.ColorRed, Ch: ' '})
}

// Tick updates the food each frame
func (food *Food) Tick(event tl.Event) {
	return
}

func randInt(max int) int {
	return int(rand.Float64() * float64(max))
}
