package arena

const (
	HumanVsAIMode = iota
	HumanVsHumanMode
	AIVsAIMode
)

type Arena struct {
	Goban
	HasWinner bool
	Players   []Player
	GameMode  int
}

func NewArena(players ...Player) *Arena {
	arena := &Arena{
		HasWinner: false,
		Players:   nil,
		GameMode:  HumanVsAIMode,
	}
	arena.Players = players
	return arena
}

var Gomoku *Arena
