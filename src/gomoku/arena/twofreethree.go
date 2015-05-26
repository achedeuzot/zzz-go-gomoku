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
	count, flankedCount := goban.countHorizontal(row, col, currentColor, 1)
	if count >= 3 && flankedCount == 0 {
		return 1
	}
	return 0
}

func (goban *Goban) checkFreeThreeVertical(row int32, col int32, currentColor int8, opponentColor int8) int {
	count, flankedCount := goban.countVertical(row, col, currentColor, 1)
	if count >= 3 && flankedCount == 0 {
		return 1
	}
	return 0
}

func (goban *Goban) checkFreeThreeDiagonal_1(row int32, col int32, currentColor int8, opponentColor int8) int {
	count, flankedCount := goban.countDiagonal1(row, col, currentColor, 1)
	if count >= 3 && flankedCount == 0 {
		return 1
	}
	return 0
}

func (goban *Goban) checkFreeThreeDiagonal_2(row int32, col int32, currentColor int8, opponentColor int8) int {
	count, flankedCount := goban.countDiagonal2(row, col, currentColor, 1)
	if count >= 3 && flankedCount == 0 {
		return 1
	}
	return 0
}
