package gui

import (
	"gomoku/arena"

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
			if x >= b.Table.pos.X+12+b.CellSize.W*int32(i) && x < b.Table.pos.X+12+b.CellSize.W*int32(i+1) &&
				y >= b.Table.pos.Y+12+b.CellSize.H*int32(j) && y < b.Table.pos.Y+12+b.CellSize.H*int32(j+1) {
				return int32(i), int32(j)
			}
		}
	}
	return -1, -1
}

func (b *Board) PlayScene() {
	Renderer.Clear()

	Renderer.Copy(b.Background.texture, &b.Background.size, &b.Background.pos)
	Renderer.Copy(b.Table.texture, &b.Table.size, &b.Table.pos)

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
			if t.Type == sdl.MOUSEBUTTONUP && t.Button == sdl.BUTTON_LEFT {
				arena.Gomoku.Goban[b.LastMousePos.X+b.LastMousePos.Y*19] = arena.WhitePlayer
			}
		}
	}

	// Display content of board in top of background
	for i := 0; i < 19; i++ {
		for j := 0; j < 19; j++ {
			currVal := arena.Gomoku.Goban[i+j*19]
			if currVal > 0 && currVal < arena.MaxGobanValue {
				Renderer.Copy(b.Pawns[currVal].texture, &b.Pawns[currVal].size,
					&sdl.Rect{
						X: b.Table.pos.X + 12 + b.Pawns[currVal].pos.W*int32(i),
						Y: b.Table.pos.Y + 12 + b.Pawns[currVal].pos.H*int32(j),
						W: b.Pawns[currVal].pos.W - 10,
						H: b.Pawns[currVal].pos.H - 10,
					})
			} else if b.LastMousePos.X == int32(i) && b.LastMousePos.Y == int32(j) {
				Renderer.Copy(b.Pawns[arena.WhitePlayer].texture, &b.Pawns[arena.WhitePlayer].size,
					&sdl.Rect{
						X: b.Table.pos.X + 12 + b.Pawns[arena.WhitePlayer].pos.W*int32(i),
						Y: b.Table.pos.Y + 12 + b.Pawns[arena.WhitePlayer].pos.H*int32(j),
						W: b.Pawns[arena.WhitePlayer].pos.W - 10,
						H: b.Pawns[arena.WhitePlayer].pos.H - 10,
					})
			}
		}
	}

	Renderer.Present()

}
