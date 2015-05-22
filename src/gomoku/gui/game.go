package gui

import (
	"gomoku/arena"
	"log"

	"github.com/veandco/go-sdl2/sdl"
)

type Game struct {
	Background   *Texture
	Table        *Texture
	Pawns        []*Texture
	CellSize     sdl.Rect
	LastMousePos sdl.Rect
}

func NewBoard() *Game {
	game := &Game{
		Background: GetTextureFromImage("data/img/bg.jpg"),
		Table:      GetTextureFromImage("data/img/board.png"),
		Pawns:      make([]*Texture, arena.MaxGobanValue),
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
	game.CellSize = sdl.Rect{X: 0, Y: 0, W: game.Table.pos.W / 19, H: game.Table.pos.H / 19}
	game.Pawns[arena.WhitePlayer].pos = game.CellSize
	game.Pawns[arena.BlackPlayer].pos = game.CellSize
	return game
}

func (g *Game) XYInCell(x int32, y int32) (int32, int32) {
	for i := 0; i < 19; i++ {
		for j := 0; j < 19; j++ {
			if x >= g.Table.pos.X+16+g.CellSize.W*int32(i) && x < g.Table.pos.X+16+g.CellSize.W*int32(i+1) &&
				y >= g.Table.pos.Y+16+g.CellSize.H*int32(j) && y < g.Table.pos.Y+16+g.CellSize.H*int32(j+1) {
				return int32(i), int32(j)
			}
		}
	}
	return -1, -1
}

func (g *Game) handleEvents() {
	for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
		switch t := event.(type) {
		case *sdl.QuitEvent:
			Running = false
		case *sdl.KeyUpEvent:
			if t.Keysym.Sym == sdl.K_ESCAPE {
				CurrScene = SceneMap["MenuMain"]
			}
		case *sdl.MouseMotionEvent:
			if i, j := g.XYInCell(t.X, t.Y); i >= 0 && j >= 0 {
				g.LastMousePos.X = i
				g.LastMousePos.Y = j
			}
		case *sdl.MouseButtonEvent:
			if isMouseButtonLeftUp(t) && isEmptyCell(g.LastMousePos.Y, g.LastMousePos.X) && arena.Gomoku.CurrPlayer.IsHuman() == true {
				// check forbidden moves
				row := g.LastMousePos.Y
				col := g.LastMousePos.X
				g.applyMove(row, col)
			}
		case *sdl.MouseWheelEvent:
			if t.Type == sdl.MOUSEWHEEL {
				arena.Gomoku.SwitchPlayers()
			}
		}
	}
}

func (g *Game) PlayScene() {
	Renderer.Clear()

	Renderer.Copy(g.Background.texture, &g.Background.size, &g.Background.pos)
	Renderer.Copy(g.Table.texture, &g.Table.size, &g.Table.pos)

	g.handleEvents()
	if arena.Gomoku.CurrPlayer.IsHuman() == false {
		row, col := arena.Gomoku.CurrPlayer.PlayMove()
		g.applyMove(row, col)
	}
	g.displayCapturedPawns()
	g.displayBoard()

	Renderer.Present()
}

func (g *Game) applyMove(row int32, col int32) {
	if isAuthorizedMove(row, col) {
		arena.Gomoku.Goban.SetElem(row, col, int8(arena.Gomoku.CurrPlayer.GetColor()))
		arena.Gomoku.Goban.Capture(row, col)
		if arena.Gomoku.Goban.IsWinningMove(row, col) {
			arena.Gomoku.CurrPlayer.SetHasWon(true)
			log.Printf("Color %d win !\n", arena.Gomoku.CurrPlayer.GetColor())
		}
		arena.Gomoku.SwitchPlayers()
	}
}

func (g *Game) displayCapturedPawns() {
	for _, player := range arena.Gomoku.Players {
		color := arena.GetOpponentColor(player.GetColor())
		var x int32
		if color == arena.WhitePlayer {
			x = DisplayMode.W / 6
		} else {
			x = (DisplayMode.W / 6) * 5
		}
		for i := player.GetCaptured(); i > 0; i-- {
			Renderer.Copy(g.Pawns[color].texture, &g.Pawns[color].size,
				&sdl.Rect{
					X: x,
					Y: DisplayMode.H - g.Pawns[color].pos.W*int32(i),
					W: g.Pawns[color].pos.W - 10,
					H: g.Pawns[color].pos.W - 10,
				})
		}
	}
}

func (g *Game) displayBoard() {
	for col := 0; col < 19; col++ {
		for row := 0; row < 19; row++ {
			currVal := arena.Gomoku.Goban.GetElem(int32(row), int32(col))
			if currVal > 0 && currVal < arena.MaxGobanValue {
				Renderer.Copy(g.Pawns[currVal].texture, &g.Pawns[currVal].size,
					&sdl.Rect{
						X: g.Table.pos.X + 16 + g.Pawns[currVal].pos.W*int32(col),
						Y: g.Table.pos.Y + 16 + g.Pawns[currVal].pos.H*int32(row),
						W: g.Pawns[currVal].pos.W - 10,
						H: g.Pawns[currVal].pos.H - 10,
					})
			} else if arena.Gomoku.CurrPlayer.IsHuman() == true && g.LastMousePos.X == int32(col) && g.LastMousePos.Y == int32(row) && isAuthorizedMove(g.LastMousePos.Y, g.LastMousePos.X) {
				Renderer.Copy(g.Pawns[arena.Gomoku.CurrPlayer.GetColor()].texture, &g.Pawns[arena.Gomoku.CurrPlayer.GetColor()].size,
					&sdl.Rect{
						X: g.Table.pos.X + 16 + g.Pawns[arena.Gomoku.CurrPlayer.GetColor()].pos.W*int32(col),
						Y: g.Table.pos.Y + 16 + g.Pawns[arena.Gomoku.CurrPlayer.GetColor()].pos.H*int32(row),
						W: g.Pawns[arena.Gomoku.CurrPlayer.GetColor()].pos.W - 10,
						H: g.Pawns[arena.Gomoku.CurrPlayer.GetColor()].pos.H - 10,
					})
			}
		}
	}
}

func isAuthorizedMove(row int32, col int32) bool {
	return arena.Gomoku.Goban.GetElem(row, col) == 0 && !arena.Gomoku.Goban.CheckTwoFreeThree(row, col, int8(arena.Gomoku.CurrPlayer.GetColor()))
}

func isEmptyCell(x int32, y int32) bool {
	if arena.Gomoku.Goban.GetElem(x, y) == 0 {
		return true
	}
	return false
}
