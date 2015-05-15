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
		Goban:     NewGoban(),
		hasWinner: false,
		gameRules: standardRules,
	}
}
