package ai

import (
	"gomoku/arena"
	"time"
)

type AI struct {
	arena.DefaultPlayer
}

func NewAI(color int) *AI {
	return &AI{
		arena.DefaultPlayer{
			Color: color,
		},
	}
}

func (ai *AI) Think(timeout time.Duration) (row int, col int) {
	// Do stuff
	return 0, 0
}

func (ai *AI) Play() (row int, col int) {
	return ai.Think(500 * time.Millisecond)
}

func (ai *AI) IsHuman() bool {
	return false
}
