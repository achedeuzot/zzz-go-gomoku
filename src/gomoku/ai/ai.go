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

func (ai *AI) think(timeout time.Duration) (row int32, col int32) {
	// Do stuff
	return 0, 0
}

func (ai *AI) PlayMove() (row int32, col int32) {
	return ai.think(500 * time.Millisecond)
}

func (ai *AI) IsHuman() bool {
	return false
}
