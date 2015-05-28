package ai

import (
	"gomoku/arena"
	"math"
	"time"
)

const (
	minimax_depth = 3
)

type AI struct {
	arena.DefaultPlayer
}

func NewAI(color int8) *AI {
	return &AI{
		arena.DefaultPlayer{
			Color:      color,
			Pawns:      0,
			TotalTurns: 0,
		},
	}
}

func (ai *AI) think() []int32 {
	move := make([]int32, 2)
	_, move = abNegamax(minimax_depth, math.Inf(-1), math.Inf(1), true)
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
	color = arena.Gomoku.ActivePlayer.GetColor()

	// Check if we’re done recursing
	if hasWon() {
		return float64(math.Inf(1)), make([]int32, 2)
	}
	if depth == 0 {
		return float64(score(color)), make([]int32, 2)
	}

	// Otherwise bubble up values from below
	bestMove := make([]int32, 2)
	bestScore := math.Inf(-1)
	for _, move := range generateNeighbors(color) {
		arena.Gomoku.Goban.SetElem(move[0], move[1], color)
		capturedPositions := arena.Gomoku.Goban.Capture(move[0], move[1])

		// Recurse
		arena.Gomoku.SwitchPlayers()
		recursedScore, _ := abNegamax(depth-1, -beta, -max(alpha, bestScore), !isMaximizer)
		currentScore := -recursedScore
		arena.Gomoku.SwitchPlayers()

		arena.Gomoku.Goban.UnCapture(capturedPositions, arena.Gomoku.OtherPlayer.GetColor())
		arena.Gomoku.Goban.SetElem(move[0], move[1], 0)
		// Update the best score
		if currentScore > bestScore {
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

func generateNeighbors(color int8) [][]int32 {
	tab := make([][]int32, 0)
	if hasPlayed() == false {
		for col := int32(7); col < 12; col++ {
			for row := int32(7); row < 12; row++ {
				if arena.Gomoku.Goban.GetElem(row, col) == 0 && !arena.Gomoku.Goban.CheckTwoFreeThree(row, col, color) {
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
			if arena.Gomoku.Goban.GetElem(row, col) == 0 && !arena.Gomoku.Goban.CheckTwoFreeThree(row, col, color) &&
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

func score(color int8) (score float64) {
	opponentColor := arena.GetOpponentColor(color)
	score = 0.0
	var col int32
	var row int32
	for col = 0; col < 19; col++ {
		for row = 0; row < 19; row++ {
			if arena.Gomoku.Goban.GetElem(row, col) == color {
				score += addCaptureScore(row, col, color)
				score += float64(arena.Gomoku.ActivePlayer.GetCaptured())
				score += addAlignedScore(row, col, color)
			} else if arena.Gomoku.Goban.GetElem(row, col) == opponentColor {
				score -= addCaptureScore(row, col, opponentColor)
				score -= float64(arena.Gomoku.OtherPlayer.GetCaptured())
				score -= addAlignedScore(row, col, opponentColor)
			}
		}
	}
	return score
}

func hasWon() bool {
	player := arena.Gomoku.ActivePlayer
	if arena.Gomoku.Goban.IsWinningState(player) == true {
		return true
	}
	return false
}

func hasPlayed() bool {
	if arena.Gomoku.ActivePlayer.GetPawns() > 0 ||
		arena.Gomoku.OtherPlayer.GetPawns() > 0 {
		return true
	}
	return false
}
