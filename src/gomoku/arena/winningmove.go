package arena

func (goban *Goban) IsWinningMove() bool {
	return goban.IsWinningState(Gomoku.ActivePlayer)
}

func (goban *Goban) IsWinningState(player Player) bool {
	if player.GetCaptured() >= 10 {
		return true
	}
	currColor := player.GetColor()
	col := int32(0)
	row := int32(0)
	for col = 0; col < 19; col++ {
		for row = 0; row < 19; row++ {
			if goban.GetElem(row, col) == currColor {
				if goban.IsWinningHorizontal(row, col, currColor) ||
					goban.IsWinningVertical(row, col, currColor) ||
					goban.IsWinningDiagonal_1(row, col, currColor) ||
					goban.IsWinningDiagonal_2(row, col, currColor) {
					return true
				}
			}
		}
	}
	return false
}

func (goban *Goban) IsWinningVertical(row int32, col int32, color int8) bool {
	count := 0
	for goban.GetTopElem(row, col) == color {
		row--
	}
	for goban.GetElem(row, col) == color {
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
	count := 0
	for goban.GetLeftElem(row, col) == color {
		col--
	}
	for goban.GetElem(row, col) == color {
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
	count := 0
	for goban.GetTopLeftElem(row, col) == color {
		col--
		row--
	}
	for goban.GetElem(row, col) == color {
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
	count := 0
	for goban.GetTopRightElem(row, col) == color {
		col++
		row--
	}
	for goban.GetElem(row, col) == color {
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
