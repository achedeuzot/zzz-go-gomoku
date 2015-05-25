package ai

import (
	"gomoku/arena"
	"log"
	// "time"
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
			Pawns: 0,
		},
	}
}

func (ai *AI) think() []int32 {
	move := make([]int32, 2)
	move[0], move[1], _ = minimax(minimax_depth, true)
	return move
}

func (ai *AI) PlayMove() (row int32, col int32) {
	ch := make(chan []int32)
	select {
	case ch <- ai.think():
		move := <-ch
		log.Println("Though about something...")
		ai.AddPawns(1)
		return move[0], move[1]
		// case <-time.After(500 * time.Millisecond):
		// 	log.Println("Though about nothing...")
		// 	return -1, -1
	}
}

func (ai *AI) IsHuman() bool {
	return false
}

func minimax(depth int, isMaximizer bool) (int32, int32, int) {
	if depth == 0 || hasWon() {
		log.Println("I'm at the end.")
		return -1, -1, score()
	}
	if isMaximizer == true {
		bestValue := -5000
		bestRow := int32(-1)
		bestCol := int32(-1)
		for idx, move := range generateNeighbors() {
			arena.Gomoku.Goban.SetElem(move[0], move[1], arena.Gomoku.CurrPlayer.GetColor())
			log.Printf("I'm trying a max neighbor [%d].\n", idx)
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
		bestValue := 5000
		bestRow := int32(-1)
		bestCol := int32(-1)
		for idx, move := range generateNeighbors() {
			arena.Gomoku.Goban.SetElem(move[0], move[1], arena.GetOpponentColor(arena.Gomoku.CurrPlayer.GetColor()))
			log.Printf("I'm trying a min neighbor [%d].\n", idx)
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
	log.Println("I'm generating a neighbor.")
	if hasPlayed() == false {
		log.Println("I'm generating a new move !")
		for col := 7; col < 12; col++ {
			for row := 7; row < 12; row++ {
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
	for col := 0; col < 19; col++ {
		for row := 0; row < 19; row++ {
			if arena.Gomoku.Goban.GetElem(int32(row), int32(col)) == 0 &&
				arena.Gomoku.Goban.IsSurounded(int32(row), int32(col)) == true {
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

func hasPlayed() bool {
	for _, player := range arena.Gomoku.Players {
		if player.GetPawns() > 0 {
			return true
		}
	}
	return false
}
