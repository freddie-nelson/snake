package main

import (
	"fmt"

	tl "github.com/JoelOtter/termloop"
)

// Score global score
var Score = 0

// ScoreText score text entity
type ScoreText struct {
	*tl.Text
}

// NewScoreText create score text
func NewScoreText(fg tl.Attr, bg tl.Attr) *ScoreText {
	return &ScoreText{tl.NewText(0, 0, fmt.Sprintf("Score: %v", Score), fg, bg)}
}

// Tick updates score text ever frame
func (text *ScoreText) Tick(_ tl.Event) {
	text.SetText(fmt.Sprintf("Score: %v", Score))
}
