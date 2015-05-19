package gui

import (
	"gomoku/arena"

	"github.com/veandco/go-sdl2/sdl"
)

type Board struct {
	Background *Texture
	Prop       *Texture
	WhitePawn  *Texture
	BlackPawn  *Texture
}

func NewBoard() *Board {
	board := &Board{
		Background: GetTextureFromImage("data/img/bg.jpg"),
		Prop:       GetTextureFromImage("data/img/board.png"),
		WhitePawn:  GetTextureFromImage("data/img/white.png"),
		BlackPawn:  GetTextureFromImage("data/img/black.png"),
	}
	// Display background to the right scale
	var ratio float64
	var finalW int32
	var finalH int32

	if board.Prop.size.W > DisplayMode.W {
		ratio = float64(DisplayMode.W) / float64(board.Prop.size.W)
		finalW = int32(float64(board.Prop.size.W) * ratio)
		finalH = int32(float64(board.Prop.size.H) * ratio)
	}

	if board.Prop.size.H > DisplayMode.H {
		ratio = float64(DisplayMode.H) / float64(board.Prop.size.H)
		finalW = int32(float64(board.Prop.size.W) * ratio)
		finalH = int32(float64(board.Prop.size.H) * ratio)
	}

	board.Background.pos = sdl.Rect{X: 0, Y: 0, W: DisplayMode.W, H: DisplayMode.H}
	board.Prop.pos = sdl.Rect{X: DisplayMode.W/2 - finalW/2, Y: 0, W: finalW, H: finalH}
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
	Renderer.Copy(b.Prop.texture, &b.Prop.size, &b.Prop.pos)

	// Display content of board in top of background
	demoArena := arena.Gomoku.Arena
	for _ = range demoArena.Goban {
		Renderer.Copy(b.WhitePawn.texture, &b.WhitePawn.size, &sdl.Rect{X: DisplayMode.W / 2, Y: 0, W: b.WhitePawn.size.W, H: b.WhitePawn.size.H})
	}
	Renderer.Present()

}
