package main

import (
	tl "github.com/JoelOtter/termloop"
)

// Snake directions
const (
	RIGHT = iota
	LEFT
	UP
	DOWN
)

// ScreenWidth the width of the screen in characters
var ScreenWidth int

// ScreenHeight the height of the screen in characters
var ScreenHeight int

// Snake the snake
type Snake struct {
	*tl.Entity
	parts     []Coord
	direction int
}

// NewSnake returns a pointer to the new snake with default values
func NewSnake() *Snake {
	// default snake parts
	parts := make([]Coord, 0)
	length := 4
	for i := length; i > 0; i-- {
		parts = append(parts, Coord{i, 1})
	}

	s := Snake{tl.NewEntity(parts[0].x, parts[0].y, 1, 1), parts, -1}

	return &s
}

func (snake *Snake) head() *Coord {
	return &snake.parts[0]
}

func (snake *Snake) tail() *Coord {
	return &snake.parts[len(snake.parts)-1]
}

// Move moves all the snake pieces forward based on the snakes direction
func (snake *Snake) Move(right int, down int) {
	head := snake.parts[0]
	head.x += right
	head.y += down

	if head.x == 0 {
		head.x = ScreenWidth - 1
	} else if head.x == ScreenWidth {
		head.x = 1
	} else if head.y == 0 {
		head.y = ScreenHeight - 1
	} else if head.y == ScreenHeight {
		head.y = 1
	}

	last := snake.parts[0]
	snake.parts[0] = head

	for i := 1; i < len(snake.parts); i++ {
		temp := snake.parts[i]
		snake.parts[i] = last
		last = temp
	}

	snake.SetPosition(snake.parts[0].x, snake.parts[0].y)
}

// Grow adds an extra part at the end of the snakes body
func (snake *Snake) Grow(amount int) {
	for i := 0; i < amount; i++ {
		tail := snake.tail()

		var new Coord
		switch snake.direction {
		case RIGHT:
			new = Coord{tail.x - 1, tail.y}
		case LEFT:
			new = Coord{tail.x + 1, tail.y}
		case UP:
			new = Coord{tail.x, tail.y - 1}
		case DOWN:
			new = Coord{tail.x, tail.y + 1}
		}

		snake.parts = append(snake.parts, new)
	}
}

// CheckCollisions checks if the snake is colliding with itself or food
func (snake *Snake) CheckCollisions(screen *tl.Screen) {
	head := snake.head()

	for i, p := range snake.parts {
		// check if snake is colliding with itself
		if p.x == head.x && p.y == head.y && i != 0 {
			GameOver = true
		}

		if p.x == TheFood.x && p.y == TheFood.y {
			TheFood.Eat()
			snake.Grow(2)
		}
	}
}

// Tick updates snake direction each frame
func (snake *Snake) Tick(event tl.Event) {
	if event.Type == tl.EventKey {
		switch event.Key {
		case tl.KeyArrowRight:
			if snake.direction != LEFT {
				snake.direction = RIGHT
			}
		case tl.KeyArrowLeft:
			if snake.direction != RIGHT {
				snake.direction = LEFT
			}
		case tl.KeyArrowUp:
			if snake.direction != DOWN {
				snake.direction = UP
			}
		case tl.KeyArrowDown:
			if snake.direction != UP {
				snake.direction = DOWN
			}
		}
	}
}

// Draw renders the snake
func (snake *Snake) Draw(screen *tl.Screen) {
	switch snake.direction {
	case RIGHT:
		snake.Move(1, 0)
	case LEFT:
		snake.Move(-1, 0)
	case UP:
		snake.Move(0, -1)
	case DOWN:
		snake.Move(0, 1)
	}

	ScreenWidth, ScreenHeight = screen.Size()
	snake.CheckCollisions(screen)

	for i, p := range snake.parts {
		if i == 0 {
			screen.RenderCell(p.x, p.y, &tl.Cell{Fg: tl.ColorBlue, Bg: tl.ColorGreen, Ch: ' '})
			continue
		}
		screen.RenderCell(p.x, p.y, &tl.Cell{Fg: tl.ColorGreen, Bg: tl.ColorGreen, Ch: ' '})
	}
}
