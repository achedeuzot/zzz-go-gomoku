package human

import (
	"gomoku/arena"
)

type Human struct {
	arena.DefaultPlayer
}

func NewHuman(color byte) *Human {
	return &Human{
		arena.DefaultPlayer{
			Color:      color,
			Pawns:      0,
			TotalTurns: 0,
		},
	}
}

func (human *Human) PlayMove() (row int32, col int32) {
	human.AddPawns(1)
	return 0, 0
}

func (human *Human) IsHuman() bool {
	return true
}
