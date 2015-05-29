package arena

func (goban *Goban) CheckTwoFreeThree(row int32, col int32, currentColor int8) bool {
	if goban.GetElem(row, col) == 0 {
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

func (goban *Goban) checkFreeThreeHorizontal(row int32, col int32, currentColor int8) int {
	ccol := col
	allowedEmptySpaces := 1
	opponentColor := GetOpponentColor(currentColor)

	count := 1
	for goban.GetLeftElem(row, col) == currentColor ||
		(goban.GetLeftElem(row, col) == 0 && allowedEmptySpaces > 0) {
		if goban.GetLeftElem(row, col) == 0 {
			allowedEmptySpaces--
		} else if goban.GetLeftElem(row, col) == currentColor {
			count++
		}
		col--
	}
	if goban.GetLeftElem(row, col) == opponentColor {
		return 0
	}
	col = ccol
	allowedEmptySpaces = 1
	for goban.GetRightElem(row, col) == currentColor ||
		(goban.GetRightElem(row, col) == 0 && allowedEmptySpaces > 0) {
		if goban.GetRightElem(row, col) == 0 {
			allowedEmptySpaces--
		} else if goban.GetRightElem(row, col) == currentColor {
			count++
		}
		col++
	}
	if goban.GetRightElem(row, col) == opponentColor {
		return 0
	}
	if count < 3 {
		return 0
	}
	return 1
}

func (goban *Goban) checkFreeThreeVertical(row int32, col int32, currentColor int8) int {
	crow := row
	allowedEmptySpaces := 1
	opponentColor := GetOpponentColor(currentColor)

	count := 1
	for goban.GetTopElem(row, col) == currentColor ||
		(goban.GetTopElem(row, col) == 0 && allowedEmptySpaces > 0) {
		if goban.GetTopElem(row, col) == 0 {
			allowedEmptySpaces--
		} else if goban.GetTopElem(row, col) == currentColor {
			count++
		}
		row--
	}
	if goban.GetTopElem(row, col) == opponentColor {
		return 0
	}
	row = crow
	allowedEmptySpaces = 1
	for goban.GetBottomElem(row, col) == currentColor ||
		(goban.GetBottomElem(row, col) == 0 && allowedEmptySpaces > 0) {
		if goban.GetBottomElem(row, col) == 0 {
			allowedEmptySpaces--
		} else if goban.GetBottomElem(row, col) == currentColor {
			count++
		}
		row++
	}
	if goban.GetBottomElem(row, col) == opponentColor {
		return 0
	}
	if count < 3 {
		return 0
	}
	return 1
}

func (goban *Goban) checkFreeThreeDiagonal_1(row int32, col int32, currentColor int8) int {
	crow := row
	ccol := col
	allowedEmptySpaces := 1
	opponentColor := GetOpponentColor(currentColor)

	count := 1
	for goban.GetTopLeftElem(row, col) == currentColor ||
		(goban.GetTopLeftElem(row, col) == 0 && allowedEmptySpaces > 0) {
		if goban.GetTopLeftElem(row, col) == 0 {
			allowedEmptySpaces--
		} else if goban.GetTopLeftElem(row, col) == currentColor {
			count++
		}
		row--
		col--
	}
	if goban.GetTopLeftElem(row, col) == opponentColor {
		return 0
	}
	row = crow
	col = ccol
	allowedEmptySpaces = 1
	for goban.GetBottomRightElem(row, col) == currentColor ||
		(goban.GetBottomRightElem(row, col) == 0 && allowedEmptySpaces > 0) {
		if goban.GetBottomRightElem(row, col) == 0 {
			allowedEmptySpaces--
		} else if goban.GetBottomRightElem(row, col) == currentColor {
			count++
		}
		row++
		col++
	}
	if goban.GetBottomRightElem(row, col) == opponentColor {
		return 0
	}
	if count < 3 {
		return 0
	}
	return 1
}

func (goban *Goban) checkFreeThreeDiagonal_2(row int32, col int32, currentColor int8) int {
	crow := row
	ccol := col
	allowedEmptySpaces := 1
	opponentColor := GetOpponentColor(currentColor)

	count := 1
	for goban.GetTopRightElem(row, col) == currentColor ||
		(goban.GetTopRightElem(row, col) == 0 && allowedEmptySpaces > 0) {
		if goban.GetTopRightElem(row, col) == 0 {
			allowedEmptySpaces--
		} else if goban.GetTopRightElem(row, col) == currentColor {
			count++
		}
		row--
		col++
	}
	if goban.GetTopRightElem(row, col) == opponentColor {
		return 0
	}
	row = crow
	col = ccol
	allowedEmptySpaces = 1
	for goban.GetBottomLeftElem(row, col) == currentColor ||
		(goban.GetBottomLeftElem(row, col) == 0 && allowedEmptySpaces > 0) {
		if goban.GetBottomLeftElem(row, col) == 0 {
			allowedEmptySpaces--
		} else if goban.GetBottomLeftElem(row, col) == currentColor {
			count++
		}
		row++
		col--
	}
	if goban.GetBottomLeftElem(row, col) == opponentColor {
		return 0
	}
	if count < 3 {
		return 0
	}
	return 1
}
