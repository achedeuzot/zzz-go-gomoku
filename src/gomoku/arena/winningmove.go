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
			if val, ok := goban.GetElem(row, col); ok == true && val == currColor {
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

func (goban *Goban) IsWinningVertical(row int32, col int32, color byte) bool {
	count := 0
	for val, ok := goban.GetTopElem(row, col); ok == true && val == color; val, ok = goban.GetTopElem(row, col) {
		row--
	}
	for val, ok := goban.GetElem(row, col); ok == true && val == color; val, ok = goban.GetElem(row, col) {
		count++
		if goban.CanBeCaptured(row, col, color) == true {
			count = 0
		}
		row++
	}
	if count >= 5 {
		return true
	}
	return false
}

func (goban *Goban) IsWinningHorizontal(row int32, col int32, color byte) bool {
	count := 0
	for val, ok := goban.GetLeftElem(row, col); ok == true && val == color; val, ok = goban.GetLeftElem(row, col) {
		col--
	}
	for val, ok := goban.GetElem(row, col); ok == true && val == color; val, ok = goban.GetElem(row, col) {
		count++
		if goban.CanBeCaptured(row, col, color) == true {
			count = 0
		}
		col++
	}
	if count >= 5 {
		return true
	}
	return false
}

func (goban *Goban) IsWinningDiagonal_1(row int32, col int32, color byte) bool {
	count := 0
	for val, ok := goban.GetTopLeftElem(row, col); ok == true && val == color; val, ok = goban.GetTopLeftElem(row, col) {
		col--
		row--
	}
	for val, ok := goban.GetElem(row, col); ok == true && val == color; val, ok = goban.GetElem(row, col) {
		count++
		if goban.CanBeCaptured(row, col, color) == true {
			count = 0
		}
		col++
		row++
	}
	if count >= 5 {
		return true
	}
	return false
}

func (goban *Goban) IsWinningDiagonal_2(row int32, col int32, color byte) bool {
	count := 0
	for val, ok := goban.GetTopRightElem(row, col); ok == true && val == color; val, ok = goban.GetTopRightElem(row, col) {
		col++
		row--
	}
	for val, ok := goban.GetElem(row, col); ok == true && val == color; val, ok = goban.GetElem(row, col) {
		count++
		if goban.CanBeCaptured(row, col, color) == true {
			count = 0
		}
		col--
		row++
	}
	if count >= 5 {
		return true
	}
	return false
}
