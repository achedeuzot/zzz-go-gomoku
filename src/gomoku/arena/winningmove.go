package arena

func (goban *Goban) IsWinningMove() bool {
	if Gomoku.CurrPlayer.GetCaptured() >= 10 {
		return true
	}
	for col := 0; col < 19; col++ {
		for row := 0; row < 19; row++ {
			currColor := Gomoku.CurrPlayer.GetColor()
			if goban.GetElem(int32(row), int32(col)) == currColor {
				if goban.IsWinningHorizontal(int32(row), int32(col)) ||
					goban.IsWinningVertical(int32(row), int32(col)) ||
					goban.IsWinningDiagonal_1(int32(row), int32(col)) ||
					goban.IsWinningDiagonal_2(int32(row), int32(col)) {
					return true
				}
			}
		}
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
		if goban.canBeCaptured(row, col, currentColor) == true {
			count = 0
		}
		count++
		row++
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
		if goban.canBeCaptured(row, col, currentColor) == true {
			count = 0
		}
		count++
		col++
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
		if goban.canBeCaptured(row, col, currentColor) == true {
			count = 0
		}
		count++
		col++
		row++
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
		if goban.canBeCaptured(row, col, currentColor) == true {
			count = 0
		}
		count++
		col--
		row++
	}
	if count >= 5 {
		return true
	}
	return false
}
