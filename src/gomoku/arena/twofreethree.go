package arena

func (goban *Goban) CheckTwoFreeThree(row int32, col int32, currentColor byte) bool {
	if val, ok := goban.GetElem(row, col); ok == true && val == 0 {
		// fake move
		goban.SetElem(row, col, currentColor)
		// count aligned elements
		totalFreeTrees := 0
		totalFreeTrees += goban.checkFreeThreeHorizontal(row, col, currentColor)
		totalFreeTrees += goban.checkFreeThreeVertical(row, col, currentColor)
		totalFreeTrees += goban.checkFreeThreeDiagonal_1(row, col, currentColor)
		totalFreeTrees += goban.checkFreeThreeDiagonal_2(row, col, currentColor)
		if totalFreeTrees >= 2 {
			goban.SetElem(row, col, 0)
			return true
		}
		// undo fake move
		goban.SetElem(row, col, 0)
	}
	return false
}

func (goban *Goban) checkFreeThreeHorizontal(row int32, col int32, currentColor byte) int {
	ccol := col
	allowedEmptySpaces := 1
	opponentColor := GetOpponentColor(currentColor)

	count := 1
	for val, ok := goban.GetLeftElem(row, col); ok == true && (val == currentColor ||
		(val == 0 && allowedEmptySpaces > 0)); val, ok = goban.GetLeftElem(row, col) {
		if val == 0 {
			allowedEmptySpaces--
		} else if val == currentColor {
			count++
		}
		col--
	}
	if val, ok := goban.GetLeftElem(row, col); ok == true && val == opponentColor {
		return 0
	}
	col = ccol
	allowedEmptySpaces = 1
	for val, ok := goban.GetRightElem(row, col); ok == true && (val == currentColor ||
		(val == 0 && allowedEmptySpaces > 0)); val, ok = goban.GetRightElem(row, col) {
		if val == 0 {
			allowedEmptySpaces--
		} else if val == currentColor {
			count++
		}
		col++
	}
	if val, ok := goban.GetRightElem(row, col); ok == true && val == opponentColor {
		return 0
	}
	if count < 3 {
		return 0
	}
	return 1
}

func (goban *Goban) checkFreeThreeVertical(row int32, col int32, currentColor byte) int {
	crow := row
	allowedEmptySpaces := 1
	opponentColor := GetOpponentColor(currentColor)

	count := 1
	for val, ok := goban.GetTopElem(row, col); ok == true && (val == currentColor ||
		(val == 0 && allowedEmptySpaces > 0)); val, ok = goban.GetTopElem(row, col) {
		if val == 0 {
			allowedEmptySpaces--
		} else if val == currentColor {
			count++
		}
		row--
	}
	if val, ok := goban.GetTopElem(row, col); ok == true && val == opponentColor {
		return 0
	}
	row = crow
	allowedEmptySpaces = 1
	for val, ok := goban.GetBottomElem(row, col); ok == true && (val == currentColor ||
		(val == 0 && allowedEmptySpaces > 0)); val, ok = goban.GetBottomElem(row, col) {
		if val == 0 {
			allowedEmptySpaces--
		} else if val == currentColor {
			count++
		}
		row++
	}
	if val, ok := goban.GetBottomElem(row, col); ok == true && val == opponentColor {
		return 0
	}
	if count < 3 {
		return 0
	}
	return 1
}

func (goban *Goban) checkFreeThreeDiagonal_1(row int32, col int32, currentColor byte) int {
	crow := row
	ccol := col
	allowedEmptySpaces := 1
	opponentColor := GetOpponentColor(currentColor)

	count := 1
	for val, ok := goban.GetTopLeftElem(row, col); ok == true && (val == currentColor ||
		(val == 0 && allowedEmptySpaces > 0)); val, ok = goban.GetTopLeftElem(row, col) {
		if val == 0 {
			allowedEmptySpaces--
		} else if val == currentColor {
			count++
		}
		row--
		col--
	}
	if val, ok := goban.GetTopLeftElem(row, col); ok == true && val == opponentColor {
		return 0
	}
	row = crow
	col = ccol
	allowedEmptySpaces = 1
	for val, ok := goban.GetBottomRightElem(row, col); ok == true && (val == currentColor ||
		(val == 0 && allowedEmptySpaces > 0)); val, ok = goban.GetBottomRightElem(row, col) {
		if val == 0 {
			allowedEmptySpaces--
		} else if val == currentColor {
			count++
		}
		row++
		col++
	}
	if val, ok := goban.GetBottomRightElem(row, col); ok == true && val == opponentColor {
		return 0
	}
	if count < 3 {
		return 0
	}
	return 1
}

func (goban *Goban) checkFreeThreeDiagonal_2(row int32, col int32, currentColor byte) int {
	crow := row
	ccol := col
	allowedEmptySpaces := 1
	opponentColor := GetOpponentColor(currentColor)

	count := 1
	for val, ok := goban.GetTopRightElem(row, col); ok == true && (val == currentColor ||
		(val == 0 && allowedEmptySpaces > 0)); val, ok = goban.GetTopRightElem(row, col) {
		if val == 0 {
			allowedEmptySpaces--
		} else if val == currentColor {
			count++
		}
		row--
		col++
	}
	if val, ok := goban.GetTopRightElem(row, col); ok == true && val == opponentColor {
		return 0
	}
	row = crow
	col = ccol
	allowedEmptySpaces = 1
	for val, ok := goban.GetBottomLeftElem(row, col); ok == true && (val == currentColor ||
		(val == 0 && allowedEmptySpaces > 0)); val, ok = goban.GetBottomLeftElem(row, col) {
		if val == 0 {
			allowedEmptySpaces--
		} else if val == currentColor {
			count++
		}
		row++
		col--
	}
	if val, ok := goban.GetBottomLeftElem(row, col); ok == true && val == opponentColor {
		return 0
	}
	if count < 3 {
		return 0
	}
	return 1
}
