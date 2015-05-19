package arena

type Arena struct {
	Goban
	hasWinner bool
}

func NewArena() *Arena {
	return &Arena{
		hasWinner: false,
	}
}

// Set goban board values
const (
	_ = iota
	ColorWhite
	ColorBlack
)

const (
	HumanVsAIMode = iota
	HumanVsHumanMode
	AIVsAIMode
)

var Gomoku struct {
	Arena    *Arena
	state    int
	gameMode int
}
