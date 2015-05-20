package human

import (
	"gomoku/arena"
	"time"
)

type Human struct {
	arena.DefaultPlayer
}

func NewHuman(color int) *Human {
	return &Human{
		arena.DefaultPlayer{
			Color: color,
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

func (human *Human) IsHuman() bool {
	return true
}
