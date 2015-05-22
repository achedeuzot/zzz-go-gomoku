package arena

func (goban *Goban) IsWinningMove(row int32, col int32) bool {
	if Gomoku.CurrPlayer.GetCaptured() >= 10 ||
		goban.IsWinningHorizontal(row, col) ||
		goban.IsWinningVertical(row, col) ||
		goban.IsWinningDiagonal_1(row, col) ||
		goban.IsWinningDiagonal_2(row, col) {
		return true
	}
	return false
}

func (goban *Goban) IsWinningVertical(row int32, col int32) bool {
	currentColor := Gomoku.CurrPlayer.GetColor()
	count := 1
	for goban.GetTopElem(row, col) == currentColor {
		row--
	}
	for goban.GetBottomElem(row, col) == currentColor {
		count++
		row++
		if goban.canBeCaptured(row, col, currentColor) {
			return false
		}
	}
	if count >= 5 {
		return true
	}
	return false
}

func (goban *Goban) IsWinningHorizontal(row int32, col int32) bool {
	currentColor := Gomoku.CurrPlayer.GetColor()
	count := 1
	for goban.GetLeftElem(row, col) == currentColor {
		col--
	}
	for goban.GetRightElem(row, col) == currentColor {
		count++
		col++
		if goban.canBeCaptured(row, col, currentColor) == true {
			return false
		}
	}
	if count >= 5 {
		return true
	}
	return false
}

func (goban *Goban) IsWinningDiagonal_1(row int32, col int32) bool {
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
		if goban.canBeCaptured(row, col, currentColor) {
			return false
		}
	}
	if count >= 5 {
		return true
	}
	return false
}

func (goban *Goban) IsWinningDiagonal_2(row int32, col int32) bool {
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
		if goban.canBeCaptured(row, col, currentColor) {
			return false
		}
	}
	if count >= 5 {
		return true
	}
	return false
}
