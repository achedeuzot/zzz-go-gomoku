package human

import (
	"gomoku/arena"
)

type Human struct {
	arena.DefaultPlayer
}

func NewHuman(color int8) *Human {
	return &Human{
		arena.DefaultPlayer{
			Color: color,
		},
	}
}

func (human *Human) PlayMove() (row int32, col int32) {
	return 0, 0
}

func (human *Human) IsHuman() bool {
	return true
}
