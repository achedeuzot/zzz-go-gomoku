package ai

import (
	"gomoku/arena"
	"math"
)

func addCaptureScore(row int32, col int32, color int8) (score float64) {
	score = 0.0
	score = float64(arena.Gomoku.Goban.CountPossibleCaptures(row, col, color))
	score = math.Pow(2.0, float64(score))
	return
}
