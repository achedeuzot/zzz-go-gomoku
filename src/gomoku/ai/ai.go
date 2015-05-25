package ai

import (
	"gomoku/arena"
	"time"
)

const (
	minimax_depth = 1
)

const (
	situ_loosing     = -5000
	situ_draw        = 0
	situ_capture_two = 500
	situ_aligned     = 1000
	situ_winning     = 5000
)

type AI struct {
	arena.DefaultPlayer
}

func NewAI(color int8) *AI {
	return &AI{
		arena.DefaultPlayer{
			Color: color,
		},
	}
}

func (ai *AI) think(timeout time.Duration) (row int32, col int32) {
	row, col, _ = minimax(minimax_depth, true)
	return
}

func (ai *AI) PlayMove() (row int32, col int32) {
	return ai.think(500 * time.Millisecond)
}

func (ai *AI) IsHuman() bool {
	return false
}

func minimax(depth int, isMaximizer bool) (int32, int32, int) {
	if depth == 0 || hasWon() {
		return -1, -1, score()
	}
	if isMaximizer == true {
		bestValue := -5000
		bestRow := int32(-1)
		bestCol := int32(-1)
		for _, moves := range generateNeighbors() {
			r, c, val := minimax(depth-1, !isMaximizer)
			if bestValue <= val {
				bestValue = val
				bestRow = r
				bestCol = c
				if r == -1 || c == -1 {
					bestRow = moves[0]
					bestCol = moves[1]
				}
			}
		}
		return bestRow, bestCol, bestValue
	} else {
		bestValue := 5000
		bestRow := int32(-1)
		bestCol := int32(-1)
		for _, moves := range generateNeighbors() {
			r, c, val := minimax(depth-1, !isMaximizer)
			if bestValue >= val {
				bestValue = val
				bestRow = r
				bestCol = c
				if r == -1 || c == -1 {
					bestRow = moves[0]
					bestCol = moves[1]
				}
			}
		}
		return bestRow, bestCol, bestValue
	}
}

func generateNeighbors() [][]int32 {
	tab := make([][]int32, 0)
	for col := 0; col < 19; col++ {
		for row := 0; row < 19; row++ {
			if arena.Gomoku.Goban.GetElem(int32(row), int32(col)) == 0 {
				move := make([]int32, 2)
				move[0] = int32(row)
				move[1] = int32(col)
				tab = append(tab, move)
			}
		}
	}
	return tab
}

func score() int {
	// heuristics moth*rfucker !
	return 1
}

func hasWon() bool {
	for _, player := range arena.Gomoku.Players {
		if player.GetHasWon() == true {
			return true
		}
	}
	return false
}
