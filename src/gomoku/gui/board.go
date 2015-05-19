package gui

import (
	"github.com/veandco/go-sdl2/sdl"
)

type Board struct {
	Background *sdl.Texture
}

func NewBoard() *Board {
	return &Board{
		Background: loadPng("data/img/board.png"),
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
	var rect sdl.Rect
	rect = sdl.Rect{X: 300, Y: 300, W: 400, H: 300}
	Renderer.SetDrawColor(255, 100, 170, 255)
	Renderer.FillRect(&rect)
	Renderer.Present()
}
