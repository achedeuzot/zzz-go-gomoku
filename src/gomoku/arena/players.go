package arena

import ()

type Player interface {
	PlayMove() (row int, col int)
	IsHuman() bool
	SetId(int)
	GetId() int
	SetColor(int8)
	GetColor() int8
	GetCaptured() int
	AddCaptured(int)
	SetHasWon(bool)
}

type DefaultPlayer struct {
	Id       int
	HasWon   bool
	Color    int8
	Captured int
}

func (dp *DefaultPlayer) SetId(newid int) {
	dp.Id = newid
}

func (dp *DefaultPlayer) GetId() int {
	return dp.Id
}

func (dp *DefaultPlayer) SetColor(color int8) {
	dp.Color = color
}

func (dp *DefaultPlayer) GetColor() int8 {
	return dp.Color
}

func (dp *DefaultPlayer) GetCaptured() int {
	return dp.Captured
}

func (dp *DefaultPlayer) AddCaptured(pawns int) {
	dp.Captured += pawns
}

func (dp *DefaultPlayer) SetHasWon(value bool) {
	dp.HasWon = value
}
