package arena

import (
	"time"
)

type Player interface {
	Think(timeout time.Duration) (row int, col int)
	Play() (row int, col int)
	IsHuman() bool
	SetId(int)
	GetId() int
	SetColor(int)
	GetColor() int
}

type DefaultPlayer struct {
	Id     int
	HasWon bool
	Color  int
}

func (dp *DefaultPlayer) SetId(newid int) {
	dp.Id = newid
}

func (dp *DefaultPlayer) GetId() int {
	return dp.Id
}

func (dp *DefaultPlayer) SetColor(color int) {
	dp.Color = color
}

func (dp *DefaultPlayer) GetColor() int {
	return dp.Color
}
