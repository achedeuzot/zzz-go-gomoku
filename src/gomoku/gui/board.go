package gui

import (
	"gomoku/arena"
	"log"

	"github.com/veandco/go-sdl2/sdl"
)

type Board struct {
	Background   *Texture
	Table        *Texture
	Pawns        []*Texture
	CellSize     sdl.Rect
	LastMousePos sdl.Rect
}

func NewBoard() *Board {
	board := &Board{
		Background: GetTextureFromImage("data/img/bg.jpg"),
		Table:      GetTextureFromImage("data/img/board.png"),
		Pawns:      make([]*Texture, arena.MaxGobanValue),
	}
	// Display background to the right scale
	var ratio float64
	var finalW int32
	var finalH int32

	if board.Table.size.W > DisplayMode.W {
		ratio = float64(DisplayMode.W) / float64(board.Table.size.W)
		finalW = int32(float64(board.Table.size.W) * ratio)
		finalH = int32(float64(board.Table.size.H) * ratio)
	}

	if board.Table.size.H > DisplayMode.H {
		ratio = float64(DisplayMode.H) / float64(board.Table.size.H)
		finalW = int32(float64(board.Table.size.W) * ratio)
		finalH = int32(float64(board.Table.size.H) * ratio)
	}

	board.Background.pos = sdl.Rect{X: 0, Y: 0, W: DisplayMode.W, H: DisplayMode.H}
	board.Table.pos = sdl.Rect{X: DisplayMode.W/2 - finalW/2, Y: 0, W: finalW, H: finalH}
	board.Pawns[arena.WhitePlayer] = GetTextureFromImage("data/img/white.png")
	board.Pawns[arena.BlackPlayer] = GetTextureFromImage("data/img/black.png")
	board.CellSize = sdl.Rect{X: 0, Y: 0, W: board.Table.pos.W / 19, H: board.Table.pos.H / 19}
	board.Pawns[arena.WhitePlayer].pos = board.CellSize
	board.Pawns[arena.BlackPlayer].pos = board.CellSize
	return board
}

func (b *Board) XYInCell(x int32, y int32) (int32, int32) {
	for i := 0; i < 19; i++ {
		for j := 0; j < 19; j++ {
			if x >= b.Table.pos.X+16+b.CellSize.W*int32(i) && x < b.Table.pos.X+16+b.CellSize.W*int32(i+1) &&
				y >= b.Table.pos.Y+16+b.CellSize.H*int32(j) && y < b.Table.pos.Y+16+b.CellSize.H*int32(j+1) {
				return int32(i), int32(j)
			}
		}
	}
	return -1, -1
}

func (b *Board) handleEvents() {
	for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
		switch t := event.(type) {
		case *sdl.QuitEvent:
			Running = false
		case *sdl.KeyUpEvent:
			if t.Keysym.Sym == sdl.K_ESCAPE {
				CurrScene = SceneMap["MenuMain"]
			}
		case *sdl.MouseMotionEvent:
			if i, j := b.XYInCell(t.X, t.Y); i >= 0 && j >= 0 {
				b.LastMousePos.X = i
				b.LastMousePos.Y = j
			}
		case *sdl.MouseButtonEvent:
			if isMouseButtonLeftUp(t) && isEmptyCell(b.LastMousePos.Y, b.LastMousePos.X) && arena.Gomoku.CurrPlayer.IsHuman() == true {
				// check forbidden moves
				row := b.LastMousePos.Y
				col := b.LastMousePos.X
				if isAuthorizedMove(row, col) {
					arena.Gomoku.Goban.SetElem(b.LastMousePos.Y, b.LastMousePos.X, int8(arena.Gomoku.CurrPlayer.GetColor()))
					arena.Gomoku.Goban.Capture(b.LastMousePos.Y, b.LastMousePos.X)
					if arena.Gomoku.Goban.CheckFiveAlign(b.LastMousePos.Y, b.LastMousePos.X) {
						arena.Gomoku.CurrPlayer.SetHasWon(true)
						log.Printf("Color %d win !\n", arena.Gomoku.CurrPlayer.GetColor())
					}
					arena.Gomoku.SwitchPlayers()
				}
			}
		case *sdl.MouseWheelEvent:
			if t.Type == sdl.MOUSEWHEEL {
				arena.Gomoku.SwitchPlayers()
			}
		}
	}
}

func (b *Board) PlayScene() {
	Renderer.Clear()

	Renderer.Copy(b.Background.texture, &b.Background.size, &b.Background.pos)
	Renderer.Copy(b.Table.texture, &b.Table.size, &b.Table.pos)

	b.handleEvents()
	b.displayCapturedPawns()
	b.displayBoard()

	Renderer.Present()
}

func (b *Board) displayCapturedPawns() {
	for _, player := range arena.Gomoku.Players {
		color := arena.GetOpponentColor(player.GetColor())
		var x int32
		if color == arena.WhitePlayer {
			x = DisplayMode.W / 6
		} else {
			x = (DisplayMode.W / 6) * 5
		}
		for i := player.GetCaptured(); i > 0; i-- {
			Renderer.Copy(b.Pawns[color].texture, &b.Pawns[color].size,
				&sdl.Rect{
					X: x,
					Y: DisplayMode.H - b.Pawns[color].pos.W*int32(i),
					W: b.Pawns[color].pos.W - 10,
					H: b.Pawns[color].pos.W - 10,
				})
		}
	}
}

func (b *Board) displayBoard() {
	for col := 0; col < 19; col++ {
		for row := 0; row < 19; row++ {
			currVal := arena.Gomoku.Goban.GetElem(int32(row), int32(col))
			if currVal > 0 && currVal < arena.MaxGobanValue {
				Renderer.Copy(b.Pawns[currVal].texture, &b.Pawns[currVal].size,
					&sdl.Rect{
						X: b.Table.pos.X + 16 + b.Pawns[currVal].pos.W*int32(col),
						Y: b.Table.pos.Y + 16 + b.Pawns[currVal].pos.H*int32(row),
						W: b.Pawns[currVal].pos.W - 10,
						H: b.Pawns[currVal].pos.H - 10,
					})
			} else if b.LastMousePos.X == int32(col) && b.LastMousePos.Y == int32(row) {
				Renderer.Copy(b.Pawns[arena.Gomoku.CurrPlayer.GetColor()].texture, &b.Pawns[arena.Gomoku.CurrPlayer.GetColor()].size,
					&sdl.Rect{
						X: b.Table.pos.X + 16 + b.Pawns[arena.Gomoku.CurrPlayer.GetColor()].pos.W*int32(col),
						Y: b.Table.pos.Y + 16 + b.Pawns[arena.Gomoku.CurrPlayer.GetColor()].pos.H*int32(row),
						W: b.Pawns[arena.Gomoku.CurrPlayer.GetColor()].pos.W - 10,
						H: b.Pawns[arena.Gomoku.CurrPlayer.GetColor()].pos.H - 10,
					})
			}
		}
	}
}

func isAuthorizedMove(row int32, col int32) bool {
	return !arena.Gomoku.Goban.CheckTwoFreeThree(row, col, int8(arena.Gomoku.CurrPlayer.GetColor()))
}

func isEmptyCell(x int32, y int32) bool {
	if arena.Gomoku.Goban.GetElem(x, y) == 0 {
		return true
	}
	return false
}
