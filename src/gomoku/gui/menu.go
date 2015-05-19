package gui

import (
	"github.com/veandco/go-sdl2/sdl"
	"github.com/veandco/go-sdl2/sdl_ttf"
)

type MenuMain struct {
	Background *Text
	Font       *ttf.Font
	Title      *Text
	Play       *Text
	Settings   *Text
	Quit       *Text
}

func NewMainMenu() *MenuMain {
	return &MenuMain{
		Background: GetTextureFromImage("data/img/bg.jpg"),
		Title:      GetTextureFromFont("data/fonts/TaiLeb.ttf", "Gogomoku", 150, sdl.Color{R: 255, G: 255, B: 255, A: 255}),
		Play:       GetTextureFromFont("data/fonts/TaiLeb.ttf", "Play", 150, sdl.Color{R: 0, G: 255, B: 255, A: 255}),
		Settings:   GetTextureFromFont("data/fonts/TaiLeb.ttf", "Settings", 150, sdl.Color{R: 0, G: 255, B: 255, A: 255}),
		Quit:       GetTextureFromFont("data/fonts/TaiLeb.ttf", "Quit", 150, sdl.Color{R: 0, G: 255, B: 255, A: 255}),
	}
}

func XYInRect(rect sdl.Rect, x int32, y int32) bool {
	return ((x > rect.X && x < rect.X+rect.W) && (y > rect.Y && y < rect.Y+rect.H))
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
		case *sdl.MouseButtonEvent:
			if t.Button == sdl.BUTTON_LEFT {
				if XYInRect(sdl.Rect{X: DisplayMode.W/2 - s.Quit.size.W/2, Y: (DisplayMode.H / 7) * 5, W: s.Quit.size.W, H: s.Quit.size.H}, t.X, t.Y) {
					Running = false
				}

			}
		}
	}
	Renderer.Clear()
	Renderer.SetDrawColor(255, 0, 0, 255)
	Renderer.FillRect(&sdl.Rect{X: 0, Y: 0, W: DisplayMode.W, H: DisplayMode.H})

	Renderer.Copy(s.Background.texture, &s.Background.size, &s.Background.size)
	Renderer.Copy(s.Title.texture, &s.Title.size, &sdl.Rect{X: DisplayMode.W/2 - s.Title.size.W/2, Y: DisplayMode.H / 7, W: s.Title.size.W, H: s.Title.size.H})
	Renderer.Copy(s.Play.texture, &s.Play.size, &sdl.Rect{X: DisplayMode.W/2 - s.Play.size.W/2, Y: (DisplayMode.H / 7) * 3, W: s.Play.size.W, H: s.Play.size.H})
	Renderer.Copy(s.Settings.texture, &s.Settings.size, &sdl.Rect{X: DisplayMode.W/2 - s.Settings.size.W/2, Y: (DisplayMode.H / 7) * 4, W: s.Settings.size.W, H: s.Settings.size.H})
	Renderer.Copy(s.Quit.texture, &s.Quit.size, &sdl.Rect{X: DisplayMode.W/2 - s.Quit.size.W/2, Y: (DisplayMode.H / 7) * 5, W: s.Quit.size.W, H: s.Quit.size.H})

	Renderer.Present()
}
