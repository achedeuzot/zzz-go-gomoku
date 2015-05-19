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
	menu := &MenuMain{
		Background: GetTextureFromImage("data/img/bg.jpg"),
		Title:      GetTextureFromFont("data/fonts/TaiLeb.ttf", "Gogomoku", 150, sdl.Color{R: 255, G: 255, B: 255, A: 255}),
		Play:       GetTextureFromFont("data/fonts/TaiLeb.ttf", "Play", 150, sdl.Color{R: 0, G: 255, B: 255, A: 255}),
		Settings:   GetTextureFromFont("data/fonts/TaiLeb.ttf", "Settings", 150, sdl.Color{R: 0, G: 255, B: 255, A: 255}),
		Quit:       GetTextureFromFont("data/fonts/TaiLeb.ttf", "Quit", 150, sdl.Color{R: 0, G: 255, B: 255, A: 255}),
	}

	menu.Background.pos = sdl.Rect{X: 0, Y: 0, W: DisplayMode.W, H: DisplayMode.H}
	menu.Title.pos = sdl.Rect{X: DisplayMode.W/2 - menu.Title.size.W/2, Y: DisplayMode.H / 7, W: menu.Title.size.W, H: menu.Title.size.H}
	menu.Play.pos = sdl.Rect{X: DisplayMode.W/2 - menu.Play.size.W/2, Y: (DisplayMode.H / 7) * 3, W: menu.Play.size.W, H: menu.Play.size.H}
	menu.Settings.pos = sdl.Rect{X: DisplayMode.W/2 - menu.Settings.size.W/2, Y: (DisplayMode.H / 7) * 4, W: menu.Settings.size.W, H: menu.Settings.size.H}
	menu.Quit.pos = sdl.Rect{X: DisplayMode.W/2 - menu.Quit.size.W/2, Y: (DisplayMode.H / 7) * 5, W: menu.Quit.size.W, H: menu.Quit.size.H}
	return menu
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
				if XYInRect(s.Quit.pos, t.X, t.Y) {
					Running = false
				} else if XYInRect(s.Play.pos, t.X, t.Y) {
					CurrScene = SceneMap["Board"]
				}

			}
		}
	}
	Renderer.Clear()
	Renderer.SetDrawColor(0, 0, 0, 255)
	Renderer.FillRect(&sdl.Rect{X: 0, Y: 0, W: DisplayMode.W, H: DisplayMode.H})

	Renderer.Copy(s.Background.texture, &s.Background.size, &s.Background.size)
	Renderer.Copy(s.Title.texture, &s.Title.size, &s.Title.pos)
	Renderer.FillRect(&s.Play.pos)
	Renderer.Copy(s.Play.texture, &s.Play.size, &s.Play.pos)
	Renderer.FillRect(&s.Settings.pos)
	Renderer.Copy(s.Settings.texture, &s.Settings.size, &s.Settings.pos)
	Renderer.FillRect(&s.Quit.pos)
	Renderer.Copy(s.Quit.texture, &s.Quit.size, &s.Quit.pos)

	Renderer.Present()
}
