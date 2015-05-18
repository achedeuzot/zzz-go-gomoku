package arena

type Arena struct {
	Goban
	hasWinner bool
	gameRules int
}

const (
	standardRules = iota
	proRules
)

func NewArena() *Arena {
	return &Arena{
		hasWinner: false,
		gameRules: standardRules,
	}
}

const (
	HumanVsAIMode = iota
	HumanVsHumanMode
	AIVsAIMode
)

var Gomoku struct {
	state    int
	gameMode int
}
