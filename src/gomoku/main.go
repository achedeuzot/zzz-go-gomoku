package main

import (
	"gomoku/ai"
	"gomoku/arena"
	"gomoku/human"
)

const (
	HumanVsAIMode = iota
	HumanVsHumanMode
	AIVsAIMode
)

var Gomoku struct {
	state    int
	gameMode int
}

func main() {
	arena.NewArena()
	ai.NewAI(0, true)
	human.NewHuman(1, false)
	return
}
