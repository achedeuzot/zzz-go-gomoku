package ai

import (
	"gomoku/arena"
	"math"
)

func addCaptureScore(row int32, col int32, color int8) (score int32) {
	score = 0
	score = arena.Gomoku.Goban.CountPossibleCaptures(row, col, color)
	score = int32(math.Pow(2.0, float64(score*2)))
	return
}
