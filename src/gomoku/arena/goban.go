package arena

// Set goban board values
const (
	_ = iota
	WhitePlayer
	BlackPlayer
	RedPawn
	CapturePawn
	MaxGobanValue
)

type Goban [361]int8

func NewGoban() *Goban {
	return &Goban{}
}

func GetOpponentColor(color int8) int8 {
	if color == WhitePlayer {
		return BlackPlayer
	}
	return WhitePlayer
}

func (goban *Goban) Copy() *Goban {
	var newgoban Goban
	for idx, val := range newgoban {
		newgoban[idx] = val
	}
	return &newgoban
}

func (goban *Goban) GetElem(row int32, col int32) int8 {
	if test := row*19 + col; test >= 0 && test <= 360 {
		return goban[row*19+col]
	}
	return -1
}

func (goban *Goban) GetTopElem(row int32, col int32) int8 {
	if row > 0 && row < 19 && col >= 0 && col < 19 {
		return goban[(row-1)*19+col]
	}
	return -1
}

func (goban *Goban) GetBottomElem(row int32, col int32) int8 {
	if row >= 0 && row < 18 && col >= 0 && col < 19 {
		return goban[(row+1)*19+col]
	}
	return -1
}

func (goban *Goban) GetLeftElem(row int32, col int32) int8 {
	if row >= 0 && row < 19 && col > 0 && col < 19 {
		return goban[row*19+(col-1)]
	}
	return -1
}

func (goban *Goban) GetRightElem(row int32, col int32) int8 {
	if row >= 0 && row < 19 && col >= 0 && col < 18 {
		return goban[row*19+(col+1)]
	}
	return -1
}

func (goban *Goban) GetTopLeftElem(row int32, col int32) int8 {
	if row > 0 && row < 19 && col > 0 && col < 19 {
		return goban[(row-1)*19+(col-1)]
	}
	return -1
}

func (goban *Goban) GetTopRightElem(row int32, col int32) int8 {
	if row > 0 && row < 19 && col >= 0 && col < 18 {
		return goban[(row-1)*19+(col+1)]
	}
	return -1
}

func (goban *Goban) GetBottomLeftElem(row int32, col int32) int8 {
	if row >= 0 && row < 18 && col > 0 && col < 19 {
		return goban[(row+1)*19+(col-1)]
	}
	return -1
}

func (goban *Goban) GetBottomRightElem(row int32, col int32) int8 {
	if row >= 0 && row < 18 && col >= 0 && col < 18 {
		return goban[(row+1)*19+(col+1)]
	}
	return -1
}

func (goban *Goban) SetElem(row int32, col int32, val int8) {
	goban[row*19+col] = val
}

func (goban *Goban) IsSurounded(row int32, col int32) bool {
	if goban.GetTopElem(row, col) > 0 ||
		goban.GetBottomElem(row, col) > 0 ||
		goban.GetLeftElem(row, col) > 0 ||
		goban.GetRightElem(row, col) > 0 ||
		goban.GetTopLeftElem(row, col) > 0 ||
		goban.GetTopRightElem(row, col) > 0 ||
		goban.GetBottomLeftElem(row, col) > 0 ||
		goban.GetBottomRightElem(row, col) > 0 {
		return true
	}
	return false
}
