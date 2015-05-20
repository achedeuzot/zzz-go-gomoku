package arena

import (
	"time"
)

type Player interface {
	SetIsWhite(bool)
	Think(timeout time.Duration) (row int, col int)
	Play() (row int, col int)
}

type DefaultPlayer struct {
	Id      int8
	IsHuman bool
	HasWon  bool
	IsWhite bool
}
