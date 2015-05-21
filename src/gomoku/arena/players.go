package arena

import ()

type Player interface {
	PlayMove() (row int, col int)
	IsHuman() bool
	SetId(int)
	GetId() int
	SetColor(int)
	GetColor() int
	GetCaptured() int
	AddCaptured(int)
	SetHasWon(bool)
}

type DefaultPlayer struct {
	Id       int
	HasWon   bool
	Color    int
	Captured int
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

func (dp *DefaultPlayer) GetCaptured() int {
	return dp.Captured
}

func (dp *DefaultPlayer) AddCaptured(pawns int) {
	dp.Captured += pawns
}

func (dp *DefaultPlayer) SetHasWon(value bool) {
	dp.HasWon = value
}
