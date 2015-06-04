package arena

type Player interface {
	PlayMove() (row int32, col int32)
	IsHuman() bool
	SetId(int)
	GetId() int
	SetColor(byte)
	GetColor() byte
	GetCaptured() int8
	AddCaptured(int8)
	GetPawns() int8
	AddPawns(int8)
	SetHasWon(bool)
	GetHasWon() bool
	AddTotalTurns()
	GetTotalTurns() int8
}

type DefaultPlayer struct {
	Id         int
	HasWon     bool
	Color      byte
	Captured   int8
	Pawns      int8
	TotalTurns int8
}

func (dp *DefaultPlayer) SetId(newid int) {
	dp.Id = newid
}

func (dp *DefaultPlayer) GetId() int {
	return dp.Id
}

func (dp *DefaultPlayer) SetColor(color byte) {
	dp.Color = color
}

func (dp *DefaultPlayer) GetColor() byte {
	return dp.Color
}

func (dp *DefaultPlayer) GetCaptured() int8 {
	return dp.Captured
}

func (dp *DefaultPlayer) AddCaptured(pawns int8) {
	dp.Captured += pawns
}

func (dp *DefaultPlayer) GetPawns() int8 {
	return dp.Pawns
}

func (dp *DefaultPlayer) AddPawns(pawns int8) {
	dp.Pawns += pawns
}

func (dp *DefaultPlayer) SetHasWon(value bool) {
	dp.HasWon = value
}

func (dp *DefaultPlayer) GetHasWon() bool {
	return dp.HasWon
}

func (dp *DefaultPlayer) AddTotalTurns() {
	dp.TotalTurns += 1
}

func (dp *DefaultPlayer) GetTotalTurns() int8 {
	return dp.TotalTurns
}
