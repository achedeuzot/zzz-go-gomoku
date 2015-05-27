package gui

import (
	"fmt"
	"gomoku/arena"

	"github.com/veandco/go-sdl2/sdl"
)

type Game struct {
	Background   *Texture
	Table        *Texture
	Pawns        []*Texture
	CellSize     sdl.Rect
	LastMousePos sdl.Rect
	Winner       []*Texture
}

func NewGame() *Game {
	game := &Game{
		Background: GetTextureFromImage("data/img/bg.jpg"),
		Table:      GetTextureFromImage("data/img/board.png"),
		Pawns:      make([]*Texture, arena.MaxGobanValue),
		Winner:     nil,
	}
	// Display background to the right scale
	var ratio float64
	var finalW int32
	var finalH int32

	if game.Table.size.W > DisplayMode.W {
		ratio = float64(DisplayMode.W) / float64(game.Table.size.W)
		finalW = int32(float64(game.Table.size.W) * ratio)
		finalH = int32(float64(game.Table.size.H) * ratio)
	}

	if game.Table.size.H > DisplayMode.H {
		ratio = float64(DisplayMode.H) / float64(game.Table.size.H)
		finalW = int32(float64(game.Table.size.W) * ratio)
		finalH = int32(float64(game.Table.size.H) * ratio)
	}

	game.Background.pos = sdl.Rect{X: 0, Y: 0, W: DisplayMode.W, H: DisplayMode.H}
	game.Table.pos = sdl.Rect{X: DisplayMode.W/2 - finalW/2, Y: 0, W: finalW, H: finalH}
	game.Pawns[arena.WhitePlayer] = GetTextureFromImage("data/img/white.png")
	game.Pawns[arena.BlackPlayer] = GetTextureFromImage("data/img/black.png")
	game.Pawns[arena.RedPawn] = GetTextureFromImage("data/img/red.png")
	game.Pawns[arena.CapturePawn] = GetTextureFromImage("data/img/capture.png")
	game.CellSize = sdl.Rect{X: 0, Y: 0, W: game.Table.pos.W / 19, H: game.Table.pos.H / 19}
	game.Pawns[arena.WhitePlayer].pos = game.CellSize
	game.Pawns[arena.BlackPlayer].pos = game.CellSize
	game.Pawns[arena.RedPawn].pos = game.CellSize
	game.Pawns[arena.CapturePawn].pos = game.CellSize
	game.Winner = make([]*Texture, 3)
	winStrWhite := fmt.Sprint("Color white won.")
	winStrBlack := fmt.Sprint("Color black won.")
	game.Winner[arena.WhitePlayer] = GetTextureFromFont(0, winStrWhite, 70, sdl.Color{R: 255, G: 255, B: 255, A: 255})
	game.Winner[arena.BlackPlayer] = GetTextureFromFont(0, winStrBlack, 70, sdl.Color{R: 255, G: 255, B: 255, A: 255})
	game.Winner[arena.WhitePlayer].pos = sdl.Rect{X: DisplayMode.W/2 - (game.Winner[arena.WhitePlayer].size.W*DisplayMode.W/2560)/2, Y: DisplayMode.H / 7, W: game.Winner[arena.WhitePlayer].size.W * DisplayMode.W / 2560, H: game.Winner[arena.WhitePlayer].size.H * DisplayMode.H / 1440}
	game.Winner[arena.BlackPlayer].pos = sdl.Rect{X: DisplayMode.W/2 - (game.Winner[arena.BlackPlayer].size.W*DisplayMode.W/2560)/2, Y: DisplayMode.H / 7, W: game.Winner[arena.BlackPlayer].size.W * DisplayMode.W / 2560, H: game.Winner[arena.WhitePlayer].size.H * DisplayMode.H / 1440}
	return game
}

func (s *Game) XYInCell(x int32, y int32) (int32, int32) {
	for i := 0; i < 19; i++ {
		for j := 0; j < 19; j++ {
			if x >= s.Table.pos.X+16+s.CellSize.W*int32(i) && x < s.Table.pos.X+16+s.CellSize.W*int32(i+1) &&
				y >= s.Table.pos.Y+16+s.CellSize.H*int32(j) && y < s.Table.pos.Y+16+s.CellSize.H*int32(j+1) {
				return int32(i), int32(j)
			}
		}
	}
	return -1, -1
}

func (s *Game) handleEvents() {
	for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
		switch t := event.(type) {
		case *sdl.QuitEvent:
			Running = false
		case *sdl.KeyUpEvent:
			if t.Keysym.Sym == sdl.K_ESCAPE {
				CurrScene = SceneMap["MenuMain"]
			}
		case *sdl.MouseMotionEvent:
			if i, j := s.XYInCell(t.X, t.Y); i >= 0 && j >= 0 {
				s.LastMousePos.X = i
				s.LastMousePos.Y = j

			}
		case *sdl.MouseButtonEvent:
			if arena.Gomoku.ActivePlayer.GetHasWon() == false && arena.Gomoku.OtherPlayer.GetHasWon() == false && isMouseButtonLeftUp(t) && isEmptyCell(s.LastMousePos.Y, s.LastMousePos.X) && arena.Gomoku.ActivePlayer.IsHuman() == true {
				// check forbidden moves
				row := s.LastMousePos.Y
				col := s.LastMousePos.X
				s.applyMove(row, col)
			}
		case *sdl.MouseWheelEvent:
			if t.Type == sdl.MOUSEWHEEL {
				arena.Gomoku.SwitchPlayers()
			}
		}
	}
}

func (s *Game) PlayScene() {
	Renderer.Clear()

	Renderer.Copy(s.Background.texture, &s.Background.size, &s.Background.pos)
	Renderer.Copy(s.Table.texture, &s.Table.size, &s.Table.pos)

	s.handleEvents()

	if arena.Gomoku.ActivePlayer.GetHasWon() == true {
		Renderer.Copy(s.Winner[arena.Gomoku.ActivePlayer.GetColor()].texture, &s.Winner[arena.Gomoku.ActivePlayer.GetColor()].size, &s.Winner[arena.Gomoku.ActivePlayer.GetColor()].pos)
	} else if arena.Gomoku.OtherPlayer.GetHasWon() == true {
		Renderer.Copy(s.Winner[arena.Gomoku.OtherPlayer.GetColor()].texture, &s.Winner[arena.Gomoku.OtherPlayer.GetColor()].size, &s.Winner[arena.Gomoku.OtherPlayer.GetColor()].pos)
	} else {
		if arena.Gomoku.ActivePlayer.IsHuman() == false {
			row, col := arena.Gomoku.ActivePlayer.PlayMove()
			s.applyMove(row, col)
		}
	}
	s.displayCapturedPawns(arena.Gomoku.ActivePlayer)
	s.displayCapturedPawns(arena.Gomoku.OtherPlayer)
	s.displayGame()

	// Display Position in top left corner
	pos := fmt.Sprintf("%2d-%2d", s.LastMousePos.Y, s.LastMousePos.X)
	postexture := GetTextureFromFont(0, pos, 70, sdl.Color{R: 255, G: 255, B: 255, A: 255})
	Renderer.Copy(postexture.texture, &postexture.size, &sdl.Rect{X: 0, Y: 0, W: postexture.size.W, H: postexture.size.H})

	Renderer.Present()
}

func (s *Game) applyMove(row int32, col int32) {
	if isAuthorizedMove(row, col) {
		arena.Gomoku.Goban.SetElem(row, col, int8(arena.Gomoku.ActivePlayer.GetColor()))
		arena.Gomoku.Goban.Capture(row, col)
		if arena.Gomoku.Goban.IsWinningMove() {
			arena.Gomoku.ActivePlayer.SetHasWon(true)
		}
		arena.Gomoku.ActivePlayer.AddPawns(1)
		arena.Gomoku.SwitchPlayers()
	}
}

func (s *Game) displayCapturedPawns(player arena.Player) {
	color := arena.GetOpponentColor(player.GetColor())
	var x int32
	if color == arena.WhitePlayer {
		x = DisplayMode.W / 6
	} else {
		x = (DisplayMode.W / 6) * 5
	}
	for i := player.GetCaptured(); i > 0; i-- {
		Renderer.Copy(s.Pawns[color].texture, &s.Pawns[color].size,
			&sdl.Rect{
				X: x,
				Y: DisplayMode.H - s.Pawns[color].pos.W*int32(i),
				W: s.Pawns[color].pos.W - 10,
				H: s.Pawns[color].pos.W - 10,
			})
	}
}

func (s *Game) displayGame() {
	for col := 0; col < 19; col++ {
		for row := 0; row < 19; row++ {
			currVal := arena.Gomoku.Goban.GetElem(int32(row), int32(col))
			if currVal > 0 && currVal < arena.MaxGobanValue {
				Renderer.Copy(s.Pawns[currVal].texture, &s.Pawns[currVal].size,
					&sdl.Rect{
						X: s.Table.pos.X + 16 + s.Pawns[currVal].pos.W*int32(col),
						Y: s.Table.pos.Y + 16 + s.Pawns[currVal].pos.H*int32(row),
						W: s.Pawns[currVal].pos.W - 10,
						H: s.Pawns[currVal].pos.H - 10,
					})
			} else if arena.Gomoku.ActivePlayer.IsHuman() == true && s.LastMousePos.X == int32(col) && s.LastMousePos.Y == int32(row) && isAuthorizedMove(s.LastMousePos.Y, s.LastMousePos.X) {
				Renderer.Copy(s.Pawns[arena.Gomoku.ActivePlayer.GetColor()].texture, &s.Pawns[arena.Gomoku.ActivePlayer.GetColor()].size,
					&sdl.Rect{
						X: s.Table.pos.X + 16 + s.Pawns[arena.Gomoku.ActivePlayer.GetColor()].pos.W*int32(col),
						Y: s.Table.pos.Y + 16 + s.Pawns[arena.Gomoku.ActivePlayer.GetColor()].pos.H*int32(row),
						W: s.Pawns[arena.Gomoku.ActivePlayer.GetColor()].pos.W - 10,
						H: s.Pawns[arena.Gomoku.ActivePlayer.GetColor()].pos.H - 10,
					})
			} else if arena.Gomoku.ActivePlayer.IsHuman() == true && currVal == 0 && !isAuthorizedMove(int32(row), int32(col)) {
				Renderer.Copy(s.Pawns[arena.RedPawn].texture, &s.Pawns[arena.RedPawn].size,
					&sdl.Rect{
						X: s.Table.pos.X + 16 + s.Pawns[arena.RedPawn].pos.W*int32(col),
						Y: s.Table.pos.Y + 16 + s.Pawns[arena.RedPawn].pos.H*int32(row),
						W: s.Pawns[arena.RedPawn].pos.W - 10,
						H: s.Pawns[arena.RedPawn].pos.H - 10,
					})
			}
			if arena.Gomoku.ActivePlayer.IsHuman() == true && currVal != 0 && arena.Gomoku.Goban.CanBeCaptured(int32(row), int32(col), currVal) {
				Renderer.Copy(s.Pawns[arena.CapturePawn].texture, &s.Pawns[arena.CapturePawn].size,
					&sdl.Rect{
						X: s.Table.pos.X + 16 + s.Pawns[arena.CapturePawn].pos.W*int32(col),
						Y: s.Table.pos.Y + 16 + s.Pawns[arena.CapturePawn].pos.H*int32(row),
						W: s.Pawns[arena.CapturePawn].pos.W - 10,
						H: s.Pawns[arena.CapturePawn].pos.H - 10,
					})
			}
		}
	}
}

func isAuthorizedMove(row int32, col int32) bool {
	return arena.Gomoku.Goban.GetElem(row, col) == 0 && !arena.Gomoku.Goban.CheckTwoFreeThree(row, col, int8(arena.Gomoku.ActivePlayer.GetColor()))
}

func isEmptyCell(x int32, y int32) bool {
	if arena.Gomoku.Goban.GetElem(x, y) == 0 {
		return true
	}
	return false
}
