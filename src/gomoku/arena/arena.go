package arena

import (
	"log"
)

const (
	HumanVsAIMode = iota
	HumanVsHumanMode
	AIVsAIMode
)

type Arena struct {
	Goban
	HasWinner  bool
	Players    []Player
	CurrPlayer Player
	GameMode   int
}

func NewArena(players ...Player) *Arena {
	if len(players) < 2 {
		log.Fatalf("Not enough players in arena: %d\n", len(players))
	}
	arena := &Arena{
		HasWinner:  false,
		Players:    make([]Player, len(players)),
		CurrPlayer: nil,
		GameMode:   HumanVsAIMode,
	}
	for idx, val := range players {
		val.SetId(idx)
		arena.Players[idx] = val
	}
	arena.CurrPlayer = arena.Players[0]
	return arena
}

var Gomoku *Arena

func (arena *Arena) SwitchPlayers() {
	nextIdx := arena.CurrPlayer.GetId() + 1
	if nextIdx >= len(arena.Players) {
		nextIdx = 0
	}
	arena.CurrPlayer = arena.Players[nextIdx]
}
