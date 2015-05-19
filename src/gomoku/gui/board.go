package gui

import (
	"gomoku/arena"

	"github.com/veandco/go-sdl2/sdl"
)

type Board struct {
	Background *Texture
	Table      *Texture
	WhitePawn  *Texture
	BlackPawn  *Texture
}

func NewBoard() *Board {
	board := &Board{
		Background: GetTextureFromImage("data/img/bg.jpg"),
		Table:      GetTextureFromImage("data/img/board.png"),
		WhitePawn:  GetTextureFromImage("data/img/white.png"),
		BlackPawn:  GetTextureFromImage("data/img/black.png"),
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
	board.WhitePawn.pos = sdl.Rect{X: 0, Y: 0, W: board.Table.pos.W / 19, H: board.Table.pos.H / 19}
	board.BlackPawn.pos = sdl.Rect{X: 0, Y: 0, W: board.Table.pos.W / 19, H: board.Table.pos.H / 19}
	return board
}

func (b *Board) PlayScene() {
	for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
		switch t := event.(type) {
		case *sdl.QuitEvent:
			Running = false
		case *sdl.KeyUpEvent:
			if t.Keysym.Sym == sdl.K_ESCAPE {
				CurrScene = SceneMap["MenuMain"]
			}
		}
	}
	Renderer.Clear()

	Renderer.Copy(b.Background.texture, &b.Background.size, &b.Background.pos)
	Renderer.Copy(b.Table.texture, &b.Table.size, &b.Table.pos)

	// Display content of board in top of background
	for i := 0; i < 19; i++ {
		for j := 0; j < 19; j++ {
			currVal := arena.Gomoku.Arena.Goban[i+j*19]
			if currVal == 1 {
				Renderer.Copy(b.WhitePawn.texture, &b.WhitePawn.size,
					&sdl.Rect{
						X: b.Table.pos.X + 12 + b.WhitePawn.pos.W*int32(i),
						Y: b.Table.pos.Y + 12 + b.WhitePawn.pos.H*int32(j),
						W: b.WhitePawn.pos.W - 10,
						H: b.WhitePawn.pos.H - 10,
					})
			} else {
				Renderer.Copy(b.BlackPawn.texture, &b.BlackPawn.size,
					&sdl.Rect{
						X: b.Table.pos.X + 12 + b.BlackPawn.pos.W*int32(i),
						Y: b.Table.pos.Y + 12 + b.BlackPawn.pos.H*int32(j),
						W: b.BlackPawn.pos.W - 10,
						H: b.BlackPawn.pos.H - 10,
					})
			}
		}
	}

	//    demoArena := arena.Gomoku.Arena
	// for idx, _ := range demoArena.Goban {

	// 	Renderer.Copy(b.WhitePawn.texture, &b.WhitePawn.size, &sdl.Rect{X: b.Table.pos.X + b.WhitePawn.pos.X*int32(idx),
	// 		Y: b.Table.pos.Y + b.WhitePawn.pos.Y*int32(idx), W: b.WhitePawn.pos.W, H: b.WhitePawn.pos.H})
	// }
	Renderer.Present()

}
