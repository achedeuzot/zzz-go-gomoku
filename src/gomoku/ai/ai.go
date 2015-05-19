package ai

import (
	"gomoku/arena"
	"time"
)

type AI struct {
	arena.DefaultPlayer
}

func NewAI(id int8, isWhite bool) *AI {
	return &AI{
		arena.DefaultPlayer{
			Id:      id,
			IsHuman: false,
			IsWhite: isWhite,
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