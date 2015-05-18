package gui

import (
	"github.com/veandco/go-sdl2/sdl"
)

type MenuMain struct {
	Background uint8
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
	var rect sdl.Rect
	rect = sdl.Rect{X: 0, Y: 0, W: 400, H: 300}
	Renderer.SetDrawColor(255, 100, 170, 255)
	Renderer.FillRect(&rect)
	Renderer.Present()
}
