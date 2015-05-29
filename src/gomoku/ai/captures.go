package ai

import "gomoku/arena"

func addCaptureScore(row int32, col int32, color int8) (score float64) {
	score = 0.0
	score = float64(arena.Gomoku.Goban.CountPossibleCaptures(row, col, color))
	score = score * score * 2
	return
}
