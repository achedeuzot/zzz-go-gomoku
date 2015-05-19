package gui

import (
	"github.com/veandco/go-sdl2/sdl"
	"github.com/veandco/go-sdl2/sdl_ttf"
)

type MenuMain struct {
	Background *Text
	Font       *ttf.Font
	Title      *Text
}

func NewMainMenu() *MenuMain {
	return &MenuMain{
		Background: GetTextureFromImage("data/img/bg.jpg"),
		Title:      GetTextureFromFont("data/fonts/TaiLeb.ttf", "Gogomoku"),
	}
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

	Renderer.Copy(s.Background.texture, &s.Background.size, &s.Background.size)
	Renderer.Copy(s.Title.texture, &s.Title.size, &s.Title.size)

	Renderer.Present()
}
