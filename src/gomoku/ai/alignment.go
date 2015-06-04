package ai

import "gomoku/arena"

func addAlignedScore(row int32, col int32, color byte) float64 {
	score := int32(0)
	currentColor := color
	score += addVerticalAlignedScore(row, col, currentColor)
	score += addHorizontalAlignedScore(row, col, currentColor)
	score += addDiagonal1AlignedScore(row, col, currentColor)
	score += addDiagonal2AlignedScore(row, col, currentColor)
	return float64(score)
}

func addVerticalAlignedScore(row int32, col int32, currentColor byte) (score int32) {
	crow := row
	flankedCount := int32(0)
	for val, ok := arena.Gomoku.Goban.GetTopElem(row, col); ok == true && val == currentColor; val, ok = arena.Gomoku.Goban.GetTopElem(row, col) {
		row--
		score += 1
	}
	if val, ok := arena.Gomoku.Goban.GetTopElem(row, col); ok == true && val != 0 {
		flankedCount++
	}
	row = crow
	for val, ok := arena.Gomoku.Goban.GetBottomElem(row, col); ok == true && val == currentColor; val, ok = arena.Gomoku.Goban.GetBottomElem(row, col) {
		row++
		score += 1
	}
	if val, ok := arena.Gomoku.Goban.GetBottomElem(row, col); ok == true && val != 0 {
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

func addHorizontalAlignedScore(row int32, col int32, currentColor byte) (score int32) {
	ccol := col
	flankedCount := int32(0)
	for val, ok := arena.Gomoku.Goban.GetLeftElem(row, col); ok == true && val == currentColor; val, ok = arena.Gomoku.Goban.GetLeftElem(row, col) {
		col--
		score += 1
	}
	if val, ok := arena.Gomoku.Goban.GetLeftElem(row, col); ok == true && val != 0 {
		flankedCount++
	}
	col = ccol
	for val, ok := arena.Gomoku.Goban.GetRightElem(row, col); ok == true && val == currentColor; val, ok = arena.Gomoku.Goban.GetRightElem(row, col) {
		row++
		score += 1
	}
	if val, ok := arena.Gomoku.Goban.GetRightElem(row, col); ok == true && val != 0 {
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

func addDiagonal1AlignedScore(row int32, col int32, currentColor byte) (score int32) {
	crow := row
	ccol := col
	flankedCount := int32(0)
	for val, ok := arena.Gomoku.Goban.GetTopLeftElem(row, col); ok == true && val == currentColor; val, ok = arena.Gomoku.Goban.GetTopLeftElem(row, col) {
		col--
		row--
		score += 1
	}
	if val, ok := arena.Gomoku.Goban.GetTopLeftElem(row, col); ok == true && val != 0 {
		flankedCount++
	}
	row = crow
	col = ccol
	for val, ok := arena.Gomoku.Goban.GetBottomRightElem(row, col); ok == true && val == currentColor; val, ok = arena.Gomoku.Goban.GetBottomRightElem(row, col) {
		row++
		col++
		score += 1
	}
	if val, ok := arena.Gomoku.Goban.GetBottomRightElem(row, col); ok == true && val != 0 {
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

func addDiagonal2AlignedScore(row int32, col int32, currentColor byte) (score int32) {
	crow := row
	ccol := col
	flankedCount := int32(0)
	for val, ok := arena.Gomoku.Goban.GetBottomLeftElem(row, col); ok == true && val == currentColor; val, ok = arena.Gomoku.Goban.GetBottomLeftElem(row, col) {
		col--
		row++
		score += 1
	}
	if val, ok := arena.Gomoku.Goban.GetBottomLeftElem(row, col); ok == true && val != 0 {
		flankedCount++
	}
	row = crow
	col = ccol
	for val, ok := arena.Gomoku.Goban.GetTopRightElem(row, col); ok == true && val == currentColor; val, ok = arena.Gomoku.Goban.GetTopRightElem(row, col) {
		row--
		col++
		score += 1
	}
	if val, ok := arena.Gomoku.Goban.GetTopRightElem(row, col); ok == true && val != 0 {
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
