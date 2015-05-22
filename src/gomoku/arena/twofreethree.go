package arena

func (goban *Goban) CheckTwoFreeThree(row int32, col int32, currentColor int8) bool {
	opponentColor := GetOpponentColor(currentColor)
	if goban.GetElem(row, col) == 0 {
		// fake move
		goban.SetElem(row, col, currentColor)
		// count aligned elements
		totalFreeTrees := 0
		totalFreeTrees += goban.checkFreeThreeHorizontal(row, col, currentColor, opponentColor)
		totalFreeTrees += goban.checkFreeThreeVertical(row, col, currentColor, opponentColor)
		totalFreeTrees += goban.checkFreeThreeDiagonal_1(row, col, currentColor, opponentColor)
		totalFreeTrees += goban.checkFreeThreeDiagonal_2(row, col, currentColor, opponentColor)
		if totalFreeTrees >= 2 {
			goban.SetElem(row, col, 0)
			return true
		}
		// undo fake move
		goban.SetElem(row, col, 0)
	}
	return false
}

func (goban *Goban) checkFreeThreeHorizontal(row int32, col int32, currentColor int8, opponentColor int8) int {
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
		return 0
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
		return 0
	}

	if count > 2 {
		return 1
	}
	return 0
}

func (goban *Goban) checkFreeThreeVertical(row int32, col int32, currentColor int8, opponentColor int8) int {
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
		return 0
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
		return 0
	}

	if count > 2 {
		return 1
	}
	return 0
}

func (goban *Goban) checkFreeThreeDiagonal_1(row int32, col int32, currentColor int8, opponentColor int8) int {
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
		return 0
	}
	allowedEmpty = true
	if goban.GetElem(row, col) == currentColor {
		count++
	}
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
		return 0
	}

	if count > 2 {
		return 1
	}
	return 0
}

func (goban *Goban) checkFreeThreeDiagonal_2(row int32, col int32, currentColor int8, opponentColor int8) int {
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
		return 0
	}
	allowedEmpty = true
	if goban.GetElem(row, col) == currentColor {
		count++
	}
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
		return 0
	}

	if count > 2 {
		return 1
	}
	return 0
}
