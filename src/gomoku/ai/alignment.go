package ai

import (
	"gomoku/arena"
	"math"
)

func addSimpleAlignedScore(row int32, col int32) (score int32) {
	score = 0
	currentColor := arena.Gomoku.CurrPlayer.GetColor()
	score += addVerticalAlignedScore(row, col, currentColor)
	score += addHorizontalAlignedScore(row, col, currentColor)
	score += addDiagonal1AlignedScore(row, col, currentColor)
	score += addDiagonal2AlignedScore(row, col, currentColor)
	opponentColor := arena.GetOpponentColor(currentColor)
	score -= addVerticalAlignedScore(row, col, opponentColor)
	score -= addHorizontalAlignedScore(row, col, opponentColor)
	score -= addDiagonal1AlignedScore(row, col, opponentColor)
	score -= addDiagonal2AlignedScore(row, col, opponentColor)
	return
}

func addAsymetricAlignedScore(row int32, col int32) (score int32) {
	score = 0
	currentColor := arena.Gomoku.CurrPlayer.GetColor()
	score += addVerticalAlignedScore(row, col, currentColor)
	score += addHorizontalAlignedScore(row, col, currentColor)
	score += addDiagonal1AlignedScore(row, col, currentColor)
	score += addDiagonal2AlignedScore(row, col, currentColor)
	opponentColor := arena.GetOpponentColor(currentColor)
	score -= 2 * addVerticalAlignedScore(row, col, opponentColor)
	score -= 2 * addHorizontalAlignedScore(row, col, opponentColor)
	score -= 2 * addDiagonal1AlignedScore(row, col, opponentColor)
	score -= 2 * addDiagonal2AlignedScore(row, col, opponentColor)
	return
}

func addVerticalAlignedScore(row int32, col int32, currentColor int8) (score int32) {
	score = 0
	crow := row
	flanked1 := false
	flanked2 := false
	for arena.Gomoku.Goban.GetTopElem(row, col) == currentColor {
		row--
		score += 1
	}
	if arena.Gomoku.Goban.GetTopElem(row, col) != 0 {
		flanked1 = true
	}
	row = crow
	for arena.Gomoku.Goban.GetBottomElem(row, col) == currentColor {
		row++
		score += 1
	}
	if arena.Gomoku.Goban.GetBottomElem(row, col) != 0 {
		flanked2 = true
	}
	if flanked1 && flanked2 {
		score = int32(math.Pow(float64(4), float64(score)))
	}
	return
}

func addHorizontalAlignedScore(row int32, col int32, currentColor int8) (score int32) {
	score = 0
	ccol := col
	flanked1 := false
	flanked2 := false
	for arena.Gomoku.Goban.GetLeftElem(row, col) == currentColor {
		col--
		score += 1
	}
	if arena.Gomoku.Goban.GetLeftElem(row, col) != 0 {
		flanked1 = true
	}
	col = ccol
	for arena.Gomoku.Goban.GetRightElem(row, col) == currentColor {
		row++
		score += 1
	}
	if arena.Gomoku.Goban.GetRightElem(row, col) != 0 {
		flanked2 = true
	}
	if flanked1 && flanked2 {
		score = int32(math.Pow(float64(4), float64(score)))
	}
	return
}

func addDiagonal1AlignedScore(row int32, col int32, currentColor int8) (score int32) {
	score = 0
	crow := row
	ccol := col
	flanked1 := false
	flanked2 := false
	for arena.Gomoku.Goban.GetTopLeftElem(row, col) == currentColor {
		col--
		row--
		score += 1
	}
	if arena.Gomoku.Goban.GetTopLeftElem(row, col) != 0 {
		flanked1 = true
	}
	row = crow
	col = ccol
	for arena.Gomoku.Goban.GetBottomRightElem(row, col) == currentColor {
		row++
		col++
		score += 1
	}
	if arena.Gomoku.Goban.GetBottomRightElem(row, col) != 0 {
		flanked2 = true
	}
	if flanked1 && flanked2 {
		score = int32(math.Pow(float64(4), float64(score)))
	}
	return
}

func addDiagonal2AlignedScore(row int32, col int32, currentColor int8) (score int32) {
	score = 0
	crow := row
	ccol := col
	flanked1 := false
	flanked2 := false
	for arena.Gomoku.Goban.GetBottomLeftElem(row, col) == currentColor {
		col--
		row++
		score += 1
	}
	if arena.Gomoku.Goban.GetBottomLeftElem(row, col) != 0 {
		flanked1 = true
	}
	row = crow
	col = ccol
	for arena.Gomoku.Goban.GetTopRightElem(row, col) == currentColor {
		row--
		col++
		score += 1
	}
	if arena.Gomoku.Goban.GetTopRightElem(row, col) != 0 {
		flanked2 = true
	}
	if flanked1 && flanked2 {
		score = int32(math.Pow(float64(4), float64(score)))
	}
	return
}
