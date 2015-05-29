package ai

import "gomoku/arena"

func addAlignedScore(row int32, col int32, color int8) float64 {
	score := int32(0)
	currentColor := color
	score += addVerticalAlignedScore(row, col, currentColor)
	score += addHorizontalAlignedScore(row, col, currentColor)
	score += addDiagonal1AlignedScore(row, col, currentColor)
	score += addDiagonal2AlignedScore(row, col, currentColor)
	return float64(score)
}

func addVerticalAlignedScore(row int32, col int32, currentColor int8) (score int32) {
	crow := row
	flankedCount := int32(0)
	for arena.Gomoku.Goban.GetTopElem(row, col) == currentColor {
		row--
		score += 1
	}
	if arena.Gomoku.Goban.GetTopElem(row, col) != 0 {
		flankedCount++
	}
	row = crow
	for arena.Gomoku.Goban.GetBottomElem(row, col) == currentColor {
		row++
		score += 1
	}
	if arena.Gomoku.Goban.GetBottomElem(row, col) != 0 {
		flankedCount++
	}
	score = score * score * score * score
	if flankedCount == 1 {
		score = score / 2
	} else if flankedCount == 2 {
		score = score / 4
	}
	return
}

func addHorizontalAlignedScore(row int32, col int32, currentColor int8) (score int32) {
	ccol := col
	flankedCount := int32(0)
	for arena.Gomoku.Goban.GetLeftElem(row, col) == currentColor {
		col--
		score += 1
	}
	if arena.Gomoku.Goban.GetLeftElem(row, col) != 0 {
		flankedCount++
	}
	col = ccol
	for arena.Gomoku.Goban.GetRightElem(row, col) == currentColor {
		row++
		score += 1
	}
	if arena.Gomoku.Goban.GetRightElem(row, col) != 0 {
		flankedCount++
	}
	score = score * score * score * score
	if flankedCount == 1 {
		score = score / 2
	} else if flankedCount == 2 {
		score = score / 4
	}
	return
}

func addDiagonal1AlignedScore(row int32, col int32, currentColor int8) (score int32) {
	crow := row
	ccol := col
	flankedCount := int32(0)
	for arena.Gomoku.Goban.GetTopLeftElem(row, col) == currentColor {
		col--
		row--
		score += 1
	}
	if arena.Gomoku.Goban.GetTopLeftElem(row, col) != 0 {
		flankedCount++
	}
	row = crow
	col = ccol
	for arena.Gomoku.Goban.GetBottomRightElem(row, col) == currentColor {
		row++
		col++
		score += 1
	}
	if arena.Gomoku.Goban.GetBottomRightElem(row, col) != 0 {
		flankedCount++
	}
	score = score * score * score * score
	if flankedCount == 1 {
		score = score / 2
	} else if flankedCount == 2 {
		score = score / 4
	}
	return
}

func addDiagonal2AlignedScore(row int32, col int32, currentColor int8) (score int32) {
	crow := row
	ccol := col
	flankedCount := int32(0)
	for arena.Gomoku.Goban.GetBottomLeftElem(row, col) == currentColor {
		col--
		row++
		score += 1
	}
	if arena.Gomoku.Goban.GetBottomLeftElem(row, col) != 0 {
		flankedCount++
	}
	row = crow
	col = ccol
	for arena.Gomoku.Goban.GetTopRightElem(row, col) == currentColor {
		row--
		col++
		score += 1
	}
	if arena.Gomoku.Goban.GetTopRightElem(row, col) != 0 {
		flankedCount++
	}
	score = score * score * score * score
	if flankedCount == 1 {
		score = score / 2
	} else if flankedCount == 2 {
		score = score / 4
	}
	return
}
