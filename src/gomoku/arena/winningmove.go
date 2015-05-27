package arena

func (goban *Goban) IsWinningMove() bool {
	return goban.IsWinningState(Gomoku.CurrPlayer)
}

func (goban *Goban) IsWinningState(player Player) bool {
	if player.GetCaptured() >= 10 {
		return true
	}
	currColor := player.GetColor()
	for col := 0; col < 19; col++ {
		for row := 0; row < 19; row++ {
			if goban.GetElem(int32(row), int32(col)) == currColor {
				if goban.IsWinningHorizontal(int32(row), int32(col), currColor) ||
					goban.IsWinningVertical(int32(row), int32(col), currColor) ||
					goban.IsWinningDiagonal_1(int32(row), int32(col), currColor) ||
					goban.IsWinningDiagonal_2(int32(row), int32(col), currColor) {
					return true
				}
			}
		}
	}
	return false
}

func (goban *Goban) IsWinningVertical(row int32, col int32, color int8) bool {
	count := 1
	for goban.GetTopElem(row, col) == color {
		row--
	}
	for goban.GetBottomElem(row, col) == color {
		if goban.CanBeCaptured(row, col, color) == true {
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

func (goban *Goban) IsWinningHorizontal(row int32, col int32, color int8) bool {
	count := 1
	for goban.GetLeftElem(row, col) == color {
		col--
	}
	for goban.GetRightElem(row, col) == color {
		if goban.CanBeCaptured(row, col, color) == true {
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

func (goban *Goban) IsWinningDiagonal_1(row int32, col int32, color int8) bool {
	count := 1
	for goban.GetTopLeftElem(row, col) == color {
		col--
		row--
	}
	for goban.GetBottomRightElem(row, col) == color {
		if goban.CanBeCaptured(row, col, color) == true {
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

func (goban *Goban) IsWinningDiagonal_2(row int32, col int32, color int8) bool {
	count := 1
	for goban.GetTopRightElem(row, col) == color {
		col++
		row--
	}
	for goban.GetBottomLeftElem(row, col) == color {
		if goban.CanBeCaptured(row, col, color) == true {
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
