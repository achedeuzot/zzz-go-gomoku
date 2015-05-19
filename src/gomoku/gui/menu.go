package gui

import (
	"github.com/veandco/go-sdl2/sdl"
)

type MenuMain struct {
	Background *sdl.Texture
}

func (s *MenuMain) PlayScene() {
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
	Renderer.SetDrawColor(255, 0, 0, 255)
	Renderer.FillRect(&sdl.Rect{X: 0, Y: 0, W: 2560, H: 1440})

	src := sdl.Rect{X: 0, Y: 0, W: 2560, H: 1440}
	dst := sdl.Rect{X: 0, Y: 0, W: 2560, H: 1440}
	Renderer.Copy(s.Background, &src, &dst)
	Renderer.Present()
}
