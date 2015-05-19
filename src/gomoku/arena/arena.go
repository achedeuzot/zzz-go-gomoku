package arena

const (
	HumanVsAIMode = iota
	HumanVsHumanMode
	AIVsAIMode
)

type Arena struct {
	Goban
	hasWinner bool
	gameMode  int
}

func NewArena() *Arena {
	return &Arena{
		hasWinner: false,
		gameMode:  HumanVsAIMode,
	}
}

var Gomoku *Arena
