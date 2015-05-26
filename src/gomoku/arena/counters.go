package arena

func (goban *Goban) countHorizontal(row int32, col int32, currentColor int8, allowedEmptySpaces int) (count int, flankedCount int) {
	ccol := col
	cAllowedEmptySpaces := allowedEmptySpaces
	opponentColor := GetOpponentColor(currentColor)

	count = 1
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
		flankedCount++
		return
	}
	col = ccol
	allowedEmptySpaces = cAllowedEmptySpaces
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
		flankedCount++
	}
	return
}

func (goban *Goban) countVertical(row int32, col int32, currentColor int8, allowedEmptySpaces int) (count int, flankedCount int) {
	crow := row
	cAllowedEmptySpaces := allowedEmptySpaces
	opponentColor := GetOpponentColor(currentColor)

	count = 1
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
		flankedCount++
		return
	}
	row = crow
	allowedEmptySpaces = cAllowedEmptySpaces
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
		flankedCount++
	}
	return
}

func (goban *Goban) countDiagonal1(row int32, col int32, currentColor int8, allowedEmptySpaces int) (count int, flankedCount int) {
	crow := row
	ccol := col
	cAllowedEmptySpaces := allowedEmptySpaces
	opponentColor := GetOpponentColor(currentColor)

	count = 1
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
		flankedCount++
		return
	}
	row = crow
	col = ccol
	allowedEmptySpaces = cAllowedEmptySpaces
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
		flankedCount++
	}
	return
}

func (goban *Goban) countDiagonal2(row int32, col int32, currentColor int8, allowedEmptySpaces int) (count int, flankedCount int) {
	crow := row
	ccol := col
	cAllowedEmptySpaces := allowedEmptySpaces
	opponentColor := GetOpponentColor(currentColor)

	count = 1
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
		flankedCount++
		return
	}
	row = crow
	col = ccol
	allowedEmptySpaces = cAllowedEmptySpaces
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
		flankedCount++
	}
	return
}
