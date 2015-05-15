package player

import (
	"time"
)

type Player struct {
	id      uint8
	isHuman bool
	hasWon  bool
	isWhite bool
}

func NewPlayer(id uint8, isHuman bool, isWhite bool) *Player {
	return &Player{
		id:      id,
		isHuman: isHuman,
		isWhite: isWhite,
	}
}

func (player *Player) Think(timeout time.Duration) (row int, col int) {
	// Do stuff
}

func (player *Player) Play() (row int, col int) {
	return player.Think(500 * time.Millisecond)
}
