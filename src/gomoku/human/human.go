package human

import (
	"gomoku/arena"
	"time"
)

type Human struct {
	arena.DefaultPlayer
}

func NewHuman(isWhite bool) *Human {
	return &Human{
		arena.DefaultPlayer{
			IsHuman: true,
			IsWhite: isWhite,
		},
	}
}

func (human *Human) Think(timeout time.Duration) (row int, col int) {
	// Do stuff
	return 0, 0
}

func (human *Human) Play() (row int, col int) {
	return human.Think(500 * time.Millisecond)
}
