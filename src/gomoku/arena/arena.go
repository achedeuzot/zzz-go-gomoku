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
