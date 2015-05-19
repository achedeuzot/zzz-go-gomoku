package gui

import (
	"github.com/veandco/go-sdl2/sdl"
)

type Board struct {
	Background *Text
}

func NewBoard() *Board {
	return &Board{
		Background: GetTextureFromImage("data/img/board.png"),
	}
}

func (b *Board) PlayScene() {
	for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
		switch t := event.(type) {
		case *sdl.QuitEvent:
			Running = false
		case *sdl.KeyUpEvent:
			if t.Keysym.Sym == sdl.K_ESCAPE {
				Running = false
			}
		}
	}
	Renderer.Clear()

	var ratio float64
	var finalW int32
	var finalH int32

	if b.Background.size.W > DisplayMode.W {
		ratio = float64(DisplayMode.W) / float64(b.Background.size.W)
		finalW = int32(float64(b.Background.size.W) * ratio)
		finalH = int32(float64(b.Background.size.H) * ratio)
	}

	if b.Background.size.H > DisplayMode.H {
		ratio = float64(DisplayMode.H) / float64(b.Background.size.H)
		finalW = int32(float64(b.Background.size.W) * ratio)
		finalH = int32(float64(b.Background.size.H) * ratio)
	}

	Renderer.Copy(b.Background.texture, &b.Background.size, &sdl.Rect{X: DisplayMode.W/2 - finalW/2, Y: 0, W: finalW, H: finalH})
	Renderer.Present()
}
