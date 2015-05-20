package ai

import (
	"gomoku/arena"
	"time"
)

type AI struct {
	arena.DefaultPlayer
}

func NewAI(isWhite bool) *AI {
	return &AI{
		arena.DefaultPlayer{
			IsHuman: false,
			IsWhite: isWhite,
		},
	}
}

func (ai *AI) SetIsWhite(state bool) {
	ai.IsWhite = state
}

func (ai *AI) Think(timeout time.Duration) (row int, col int) {
	// Do stuff
	return 0, 0
}

func (ai *AI) Play() (row int, col int) {
	return ai.Think(500 * time.Millisecond)
}
