package arena

// Set goban board values
const (
	_ = iota
	WhitePlayer
	BlackPlayer
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

func (goban *Goban) Capture(row int32, col int32) {
	currentColor := Gomoku.CurrPlayer.GetColor()
	opponentColor := GetOpponentColor(Gomoku.CurrPlayer.GetColor())
	if goban.canCaptureUp(row, col, currentColor, opponentColor) {
		goban.SetElem(row-1, col, 0)
		goban.SetElem(row-2, col, 0)
		Gomoku.CurrPlayer.AddCaptured(2)
	}
}

func (goban *Goban) canCaptureUp(row int32, col int32, currentColor int8, opponentColor int8) bool {
	for idx := 0; idx < 3; idx++ {
		topElem := goban.GetTopElem(row, col)
		if topElem == 0 {
			return false
		} else if idx < 2 && topElem != opponentColor {
			return false
		} else if idx == 2 && topElem != currentColor {
			return false
		}
		row--
	}
	return true
}

func (goban *Goban) CheckTwoFreeThree(row int32, col int32, currentColor int8) bool {
	opponentColor := GetOpponentColor(currentColor)
	if goban.GetElem(row, col) == 0 {
		// fake move
		goban.SetElem(row, col, currentColor)
		// count aligned elements
		totalFreeTrees := 0
		if goban.CheckFreeThreeHorizontal(row, col, currentColor, opponentColor) {
			totalFreeTrees++
		}
		if goban.CheckFreeThreeVertical(row, col, currentColor, opponentColor) {
			totalFreeTrees++
		}
		if goban.CheckFreeThreeDiagnoal_1(row, col, currentColor, opponentColor) {
			totalFreeTrees++
		}
		if goban.CheckFreeThreeDiagnoal_2(row, col, currentColor, opponentColor) {
			totalFreeTrees++
		}
		if totalFreeTrees >= 2 {
			goban.SetElem(row, col, 0)
			return true
		}
		// undo fake move
		goban.SetElem(row, col, 0)
	}
	return false
}

func (goban *Goban) CheckFreeThreeHorizontal(row int32, col int32, currentColor int8, opponentColor int8) bool {
	count := 0
	allowedEmpty := true

	for goban.GetLeftElem(row, col) == currentColor ||
		(goban.GetLeftElem(row, col) == 0 && allowedEmpty == true) {
		if goban.GetLeftElem(row, col) == 0 {
			allowedEmpty = false
		}
		col--
	}
	if goban.GetLeftElem(row, col) == opponentColor {
		return false
	}
	allowedEmpty = true
	for goban.GetRightElem(row, col) == currentColor ||
		(goban.GetRightElem(row, col) == 0 && allowedEmpty == true) {
		if goban.GetRightElem(row, col) == 0 {
			allowedEmpty = false
		} else if goban.GetRightElem(row, col) == currentColor {
			count++
		}
		col++
	}
	if goban.GetRightElem(row, col) == opponentColor {
		return false
	}

	if count > 2 {
		return true
	}
	return false
}

func (goban *Goban) CheckFreeThreeVertical(row int32, col int32, currentColor int8, opponentColor int8) bool {
	count := 0
	allowedEmpty := true

	for goban.GetTopElem(row, col) == currentColor ||
		(goban.GetTopElem(row, col) == 0 && allowedEmpty == true) {
		if goban.GetTopElem(row, col) == 0 {
			allowedEmpty = false
		}
		row--
	}
	if goban.GetTopElem(row, col) == opponentColor {
		return false
	}
	allowedEmpty = true
	for goban.GetBottomElem(row, col) == currentColor ||
		(goban.GetBottomElem(row, col) == 0 && allowedEmpty == true) {
		if goban.GetBottomElem(row, col) == 0 {
			allowedEmpty = false
		} else if goban.GetBottomElem(row, col) == currentColor {
			count++
		}
		row++
	}
	if goban.GetBottomElem(row, col) == opponentColor {
		return false
	}

	if count > 2 {
		return true
	}
	return false
}

func (goban *Goban) CheckFreeThreeDiagnoal_1(row int32, col int32, currentColor int8, opponentColor int8) bool {
	count := 0
	allowedEmpty := true

	for goban.GetTopLeftElem(row, col) == currentColor ||
		(goban.GetTopLeftElem(row, col) == 0 && allowedEmpty == true) {
		if goban.GetTopLeftElem(row, col) == 0 {
			allowedEmpty = false
		}
		col--
		row--
	}
	if goban.GetTopLeftElem(row, col) == opponentColor {
		return false
	}
	allowedEmpty = true
	for goban.GetBottomRightElem(row, col) == currentColor ||
		(goban.GetBottomRightElem(row, col) == 0 && allowedEmpty == true) {
		if goban.GetBottomRightElem(row, col) == 0 {
			allowedEmpty = false
		} else if goban.GetBottomRightElem(row, col) == currentColor {
			count++
		}
		col++
		row++
	}
	if goban.GetBottomRightElem(row, col) == opponentColor {
		return false
	}

	if count > 2 {
		return true
	}
	return false
}

func (goban *Goban) CheckFreeThreeDiagnoal_2(row int32, col int32, currentColor int8, opponentColor int8) bool {
	count := 0
	allowedEmpty := true

	for goban.GetTopRightElem(row, col) == currentColor ||
		(goban.GetTopRightElem(row, col) == 0 && allowedEmpty == true) {
		if goban.GetTopRightElem(row, col) == 0 {
			allowedEmpty = false
		}
		col++
		row--
	}
	if goban.GetTopRightElem(row, col) == opponentColor {
		return false
	}
	allowedEmpty = true
	for goban.GetBottomLeftElem(row, col) == currentColor ||
		(goban.GetBottomLeftElem(row, col) == 0 && allowedEmpty == true) {
		if goban.GetBottomLeftElem(row, col) == 0 {
			allowedEmpty = false
		} else if goban.GetBottomLeftElem(row, col) == currentColor {
			count++
		}
		col--
		row++
	}
	if goban.GetBottomLeftElem(row, col) == opponentColor {
		return false
	}

	if count > 2 {
		return true
	}
	return false
}

func (goban *Goban) CheckFiveAlign(row int32, col int32) bool {
	if goban.CheckFiveAlignHorizontal(row, col) ||
		goban.CheckFiveAlignVertical(row, col) ||
		goban.CheckFiveAlignDiagonal_1(row, col) ||
		goban.CheckFiveAlignDiagonal_2(row, col) {
		return true
	}
	return false
}

func (goban *Goban) CheckFiveAlignVertical(row int32, col int32) bool {
	currentColor := Gomoku.CurrPlayer.GetColor()
	count := 1
	for goban.GetTopElem(row, col) == currentColor {
		row--
	}
	for goban.GetBottomElem(row, col) == currentColor {
		count++
		row++
	}
	if count >= 5 {
		return true
	}
	return false
}

func (goban *Goban) CheckFiveAlignHorizontal(row int32, col int32) bool {
	currentColor := Gomoku.CurrPlayer.GetColor()
	count := 1
	for goban.GetLeftElem(row, col) == currentColor {
		col--
	}
	for goban.GetRightElem(row, col) == currentColor {
		count++
		col++
	}
	if count >= 5 {
		return true
	}
	return false
}

func (goban *Goban) CheckFiveAlignDiagonal_1(row int32, col int32) bool {
	currentColor := Gomoku.CurrPlayer.GetColor()
	count := 1
	for goban.GetTopLeftElem(row, col) == currentColor {
		col--
		row--
	}
	for goban.GetBottomRightElem(row, col) == currentColor {
		count++
		col++
		row++
	}
	if count >= 5 {
		return true
	}
	return false
}

func (goban *Goban) CheckFiveAlignDiagonal_2(row int32, col int32) bool {
	currentColor := Gomoku.CurrPlayer.GetColor()
	count := 1
	for goban.GetTopRightElem(row, col) == currentColor {
		col++
		row--
	}
	for goban.GetBottomLeftElem(row, col) == currentColor {
		count++
		col--
		row++
	}
	if count >= 5 {
		return true
	}
	return false
}

func (goban *Goban) GetElem(row int32, col int32) int8 {
	if row*19+col > 361 || row*19+col < 0 {
		return -1
	}
	return goban[row*19+col]
}

func (goban *Goban) GetTopElem(row int32, col int32) int8 {
	if row > 0 && row < 19 && col > 0 && col < 19 {
		return goban[(row-1)*19+col]
	}
	return -1
}

func (goban *Goban) GetBottomElem(row int32, col int32) int8 {
	if row > 0 && row < 18 && col > 0 && col < 19 {
		return goban[(row+1)*19+col]
	}
	return -1
}

func (goban *Goban) GetLeftElem(row int32, col int32) int8 {
	if row > 0 && row < 19 && col > 0 && col < 19 {
		return goban[row*19+(col-1)]
	}
	return -1
}

func (goban *Goban) GetRightElem(row int32, col int32) int8 {
	if row > 0 && row < 19 && col > 0 && col < 18 {
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
	if row > 0 && row < 19 && col > 0 && col < 18 {
		return goban[(row-1)*19+(col+1)]
	}
	return -1
}

func (goban *Goban) GetBottomLeftElem(row int32, col int32) int8 {
	if row > 0 && row < 18 && col > 0 && col < 19 {
		return goban[(row+1)*19+(col-1)]
	}
	return -1
}

func (goban *Goban) GetBottomRightElem(row int32, col int32) int8 {
	if row > 0 && row < 18 && col > 0 && col < 18 {
		return goban[(row+1)*19+(col+1)]
	}
	return -1
}

func (goban *Goban) SetElem(row int32, col int32, val int8) {
	goban[row*19+col] = val
}
