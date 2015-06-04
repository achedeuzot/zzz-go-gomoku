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

type Goban [361]byte

func NewGoban() *Goban {
	return &Goban{}
}

func GetOpponentColor(color byte) byte {
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

func (goban *Goban) GetHash() string {
	return string(goban[:])
}

func (goban *Goban) GetElem(row int32, col int32) (byte, bool) {
	if test := row*19 + col; test >= 0 && test <= 360 {
		return goban[row*19+col], true
	}
	return 0, false
}

func (goban *Goban) GetTopElem(row int32, col int32) (byte, bool) {
	if row > 0 && row < 19 && col >= 0 && col < 19 {
		return goban[(row-1)*19+col], true
	}
	return 0, false
}

func (goban *Goban) GetBottomElem(row int32, col int32) (byte, bool) {
	if row >= 0 && row < 18 && col >= 0 && col < 19 {
		return goban[(row+1)*19+col], true
	}
	return 0, false
}

func (goban *Goban) GetLeftElem(row int32, col int32) (byte, bool) {
	if row >= 0 && row < 19 && col > 0 && col < 19 {
		return goban[row*19+(col-1)], true
	}
	return 0, false
}

func (goban *Goban) GetRightElem(row int32, col int32) (byte, bool) {
	if row >= 0 && row < 19 && col >= 0 && col < 18 {
		return goban[row*19+(col+1)], true
	}
	return 0, false
}

func (goban *Goban) GetTopLeftElem(row int32, col int32) (byte, bool) {
	if row > 0 && row < 19 && col > 0 && col < 19 {
		return goban[(row-1)*19+(col-1)], true
	}
	return 0, false
}

func (goban *Goban) GetTopRightElem(row int32, col int32) (byte, bool) {
	if row > 0 && row < 19 && col >= 0 && col < 18 {
		return goban[(row-1)*19+(col+1)], true
	}
	return 0, false
}

func (goban *Goban) GetBottomLeftElem(row int32, col int32) (byte, bool) {
	if row >= 0 && row < 18 && col > 0 && col < 19 {
		return goban[(row+1)*19+(col-1)], true
	}
	return 0, false
}

func (goban *Goban) GetBottomRightElem(row int32, col int32) (byte, bool) {
	if row >= 0 && row < 18 && col >= 0 && col < 18 {
		return goban[(row+1)*19+(col+1)], true
	}
	return 0, false
}

func (goban *Goban) SetElem(row int32, col int32, val byte) {
	goban[row*19+col] = val
}

func (goban *Goban) IsSurounded(row int32, col int32) bool {
	if val, ok := goban.GetTopElem(row, col); ok == true && val > 0 {
		return true
	}
	if val, ok := goban.GetBottomElem(row, col); ok == true && val > 0 {
		return true
	}
	if val, ok := goban.GetLeftElem(row, col); ok == true && val > 0 {
		return true
	}
	if val, ok := goban.GetRightElem(row, col); ok == true && val > 0 {
		return true
	}
	if val, ok := goban.GetTopLeftElem(row, col); ok == true && val > 0 {
		return true
	}
	if val, ok := goban.GetTopRightElem(row, col); ok == true && val > 0 {
		return true
	}
	if val, ok := goban.GetBottomLeftElem(row, col); ok == true && val > 0 {
		return true
	}
	if val, ok := goban.GetBottomRightElem(row, col); ok == true && val > 0 {
		return true
	}
	return false
}
