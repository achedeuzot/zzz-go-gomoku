package ai

import (
	"gomoku/arena"
	"time"
)

type AI struct {
	arena.DefaultPlayer
}

func NewAI(color int8) *AI {
	return &AI{
		arena.DefaultPlayer{
			Color: color,
		},
	}
}

func (ai *AI) think(timeout time.Duration) (row int, col int) {
	// Do stuff
	return 0, 0
}

func (ai *AI) PlayMove() (row int, col int) {
	return ai.think(500 * time.Millisecond)
}

func (ai *AI) IsHuman() bool {
	return false
}
