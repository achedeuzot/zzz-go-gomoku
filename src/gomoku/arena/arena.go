package arena

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

func NewArena(firstPlayer Player, players ...Player) *Arena {
	firstPlayer.SetIsWhite(false) // first player is always black
	arena := &Arena{
		HasWinner:  false,
		Players:    nil,
		CurrPlayer: firstPlayer,
		GameMode:   HumanVsAIMode,
	}
	arena.Players = players
	return arena
}

var Gomoku *Arena
