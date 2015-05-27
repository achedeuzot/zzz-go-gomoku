package arena

type Arena struct {
	Goban
	ActivePlayer Player
	OtherPlayer  Player
}

func NewArena(firstPlayer, secondPlayer Player) *Arena {
	arena := &Arena{
		ActivePlayer: firstPlayer,
		OtherPlayer:  secondPlayer,
	}
	return arena
}

var Gomoku *Arena

func (arena *Arena) SwitchPlayers() {
	arena.ActivePlayer, arena.OtherPlayer = arena.OtherPlayer, arena.ActivePlayer
}
