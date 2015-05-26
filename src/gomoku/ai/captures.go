package ai

import (
	"gomoku/arena"
	"math"
)

func addCaptureScore(row int32, col int32) (score int32) {
	score = 0
	currentColor := arena.Gomoku.CurrPlayer.GetColor()
	currScore := arena.Gomoku.Goban.CountPossibleCaptures(row, col, currentColor)
	currScore = int32(math.Pow(float64(3), float64(currScore)))
	opponentColor := arena.GetOpponentColor(currentColor)
	opponentScore := arena.Gomoku.Goban.CountPossibleCaptures(row, col, opponentColor)
	opponentScore = int32(math.Pow(float64(3), float64(opponentScore)))
	score = (2 * opponentScore) - currScore
	return
}
