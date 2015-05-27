package ai

import (
	"gomoku/arena"
	"log"
	"math"
	"time"
)

const (
	minimax_depth = 2
)

type AI struct {
	arena.DefaultPlayer
}

func NewAI(color int8) *AI {
	return &AI{
		arena.DefaultPlayer{
			Color: color,
			Pawns: 0,
		},
	}
}

func (ai *AI) think() []int32 {
	move := make([]int32, 2)
	_, move = abNegamax(minimax_depth, math.Inf(-1), math.Inf(1), true)
	// move[0], move[1], _ = minimax(minimax_depth, true)
	return move
}

func (ai *AI) PlayMove() (row int32, col int32) {
	ch := make(chan []int32, 1)
	select {
	case ch <- ai.think():
		move := <-ch
		return move[0], move[1]
	case <-time.After(500 * time.Millisecond):
		return -1, -1
	}
}

func (ai *AI) IsHuman() bool {
	return false
}

func abNegamax(depth int, alpha float64, beta float64, isMaximizer bool) (float64, []int32) {
	var color int8
	if isMaximizer == true {
		color = arena.Gomoku.CurrPlayer.GetColor()
	} else {
		color = arena.GetOpponentColor(arena.Gomoku.CurrPlayer.GetColor())
	}

	// Check if we’re done recursing
	if hasWon(isMaximizer) {
		log.Printf("Detect Winning State at depth %d\n", depth)
		return float64(math.Inf(1)), make([]int32, 2)
	}
	if depth == 0 {
		return float64(score()), make([]int32, 2)
	}

	// Otherwise bubble up values from below
	bestMove := make([]int32, 2)
	bestScore := math.Inf(-1)
	for _, move := range generateNeighbors() {
		arena.Gomoku.Goban.SetElem(move[0], move[1], color)

		// Recurse
		recursedScore, _ := abNegamax(depth-1, -beta, -max(alpha, bestScore), !isMaximizer)
		currentScore := -recursedScore

		log.Printf("[%d] Set elem on %2d-%2d - recursedScore: %f | currentScore: %f\n", depth, move[0], move[1], recursedScore, currentScore)

		arena.Gomoku.Goban.SetElem(move[0], move[1], 0)
		// Update the best score
		if currentScore > bestScore {
			log.Println("currentScore > bestScore")
			bestScore = currentScore
			bestMove = move

			// If we're outside the bounds, then prune: exit now !
			if bestScore >= beta {
				return bestScore, bestMove
			}
		}
	}
	return bestScore, bestMove
}

func max(a, b float64) float64 {
	if a >= b {
		return a
	}
	return b
}

func minimax(depth int, isMaximizer bool) (int32, int32, int32) {
	if hasWon(isMaximizer) {
		return -1, -1, 999999999
	}
	if depth == 0 {
		return -1, -1, score()
	}
	if isMaximizer == true {
		var bestValue int32 = -999999999
		bestRow := int32(-1)
		bestCol := int32(-1)
		for _, move := range generateNeighbors() {
			arena.Gomoku.Goban.SetElem(move[0], move[1], arena.Gomoku.CurrPlayer.GetColor())
			r, c, val := minimax(depth-1, !isMaximizer)
			if bestValue <= val {
				bestValue = val
				bestRow = r
				bestCol = c
				if r == -1 || c == -1 {
					bestRow = move[0]
					bestCol = move[1]
				}
			}
			arena.Gomoku.Goban.SetElem(move[0], move[1], 0)
		}
		return bestRow, bestCol, bestValue
	} else {
		var bestValue int32 = 999999999
		bestRow := int32(-1)
		bestCol := int32(-1)
		for _, move := range generateNeighbors() {
			arena.Gomoku.Goban.SetElem(move[0], move[1], arena.GetOpponentColor(arena.Gomoku.CurrPlayer.GetColor()))
			r, c, val := minimax(depth-1, !isMaximizer)
			if bestValue >= val {
				bestValue = val
				bestRow = r
				bestCol = c
				if r == -1 || c == -1 {
					bestRow = move[0]
					bestCol = move[1]
				}
			}
			arena.Gomoku.Goban.SetElem(move[0], move[1], 0)
		}
		return bestRow, bestCol, bestValue
	}
}

func generateNeighbors() [][]int32 {
	tab := make([][]int32, 0)
	if hasPlayed() == false {
		for col := int32(7); col < 12; col++ {
			for row := int32(7); row < 12; row++ {
				if arena.Gomoku.Goban.GetElem(row, col) == 0 {
					move := make([]int32, 2)
					move[0] = row
					move[1] = col
					tab = append(tab, move)
				}
			}
		}
		return tab
	}
	for col := int32(0); col < 19; col++ {
		for row := int32(0); row < 19; row++ {
			if arena.Gomoku.Goban.GetElem(row, col) == 0 &&
				arena.Gomoku.Goban.IsSurounded(row, col) == true {
				move := make([]int32, 2)
				move[0] = row
				move[1] = col
				tab = append(tab, move)
			}
		}
	}
	return tab
}

func score() (score int32) {
	score = 0
	var col int32
	var row int32
	for col = 0; col < 19; col++ {
		for row = 0; row < 19; row++ {
			if arena.Gomoku.Goban.GetElem(row, col) != 0 {
				score += addCaptureScore(row, col)
				score += addAsymetricAlignedScore(row, col)
			}
		}
	}
	return score
}

func hasWon(isMaximizer bool) bool {
	var player arena.Player
	if isMaximizer == true {
		player = arena.Gomoku.CurrPlayer
	} else {
		for _, tmpPlayer := range arena.Gomoku.Players {
			if tmpPlayer != arena.Gomoku.CurrPlayer {
				player = tmpPlayer
				break
			}
		}
	}
	if arena.Gomoku.Goban.IsWinningState(player) == true {
		return true
	}
	return false
}

func hasPlayed() bool {
	for _, player := range arena.Gomoku.Players {
		if player.GetPawns() > 0 {
			return true
		}
	}
	return false
}
