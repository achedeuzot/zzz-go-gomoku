package gui

import (
	"gomoku/arena"
	"gomoku/human"
	"gomoku/ai"

	"github.com/veandco/go-sdl2/sdl"
	"github.com/veandco/go-sdl2/sdl_ttf"
)

type MenuMain struct {
	Background *Texture
	Font       *ttf.Font
	Title      *Texture
	Play       *Texture
	Settings   *Texture
	Quit       *Texture
}

func NewMainMenu() *MenuMain {
	menu := &MenuMain{
		Background: GetTextureFromImage("data/img/bg.jpg"),
		Title:      GetTextureFromFont("data/fonts/TaiLeb.ttf", "Gogomoku", 150, sdl.Color{R: 255, G: 255, B: 255, A: 255}),
		Play:       GetTextureFromImage("data/img/button_play.png"),
		Settings:   GetTextureFromImage("data/img/button_options.png"),
		Quit:       GetTextureFromImage("data/img/button_exit.png"),
	}

	menu.Background.pos = sdl.Rect{X: 0, Y: 0, W: DisplayMode.W, H: DisplayMode.H}
	menu.Title.pos = sdl.Rect{X: DisplayMode.W/2 - (menu.Title.size.W*DisplayMode.W/2560)/2, Y: DisplayMode.H / 7, W: menu.Title.size.W * DisplayMode.W / 2560, H: menu.Title.size.H * DisplayMode.H / 1440}
	menu.Play.pos = sdl.Rect{X: DisplayMode.W/2 - (menu.Play.size.W*DisplayMode.W/2560)/2, Y: (DisplayMode.H / 7) * 3, W: menu.Play.size.W * DisplayMode.W / 2560, H: menu.Play.size.H * DisplayMode.H / 1440}
	menu.Settings.pos = sdl.Rect{X: DisplayMode.W/2 - (menu.Settings.size.W*DisplayMode.W/2560)/2, Y: (DisplayMode.H / 7) * 4, W: menu.Settings.size.W * DisplayMode.W / 2560, H: menu.Settings.size.H * DisplayMode.H / 1440}
	menu.Quit.pos = sdl.Rect{X: DisplayMode.W/2 - (menu.Quit.size.W*DisplayMode.W/2560)/2, Y: (DisplayMode.H / 7) * 5, W: menu.Quit.size.W * DisplayMode.W / 2560, H: menu.Quit.size.H * DisplayMode.H / 1440}
	return menu
}

func XYInRect(rect sdl.Rect, x int32, y int32) bool {
	return ((x > rect.X && x < rect.X+rect.W) && (y > rect.Y && y < rect.Y+rect.H))
}

func (s *MenuMain) handleEvents() {
	for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
		switch t := event.(type) {
		case *sdl.QuitEvent:
			Running = false
		case *sdl.KeyUpEvent:
			if t.Keysym.Sym == sdl.K_ESCAPE {
				Running = false
			}
		case *sdl.MouseButtonEvent:
			if isMouseButtonLeftUp(t) {
				if XYInRect(s.Quit.pos, t.X, t.Y) {
					Running = false
					break
				} else if XYInRect(s.Play.pos, t.X, t.Y) {
					humanPlayer := human.NewHuman(arena.BlackPlayer)
					aiPlayer := ai.NewAI(arena.WhitePlayer)
					arena.Gomoku = arena.NewArena(humanPlayer, aiPlayer)
					CurrScene = SceneMap["Board"]
					break
				}

			}
		}
	}
}

func (s *MenuMain) PlayScene() {
	s.handleEvents()

	Renderer.Clear()
	Renderer.SetDrawColor(0, 0, 0, 255)
	Renderer.FillRect(&sdl.Rect{X: 0, Y: 0, W: DisplayMode.W, H: DisplayMode.H})

	Renderer.Copy(s.Background.texture, &s.Background.size, &s.Background.pos)
	Renderer.Copy(s.Title.texture, &s.Title.size, &s.Title.pos)
	Renderer.Copy(s.Play.texture, &s.Play.size, &s.Play.pos)
	Renderer.Copy(s.Settings.texture, &s.Settings.size, &s.Settings.pos)
	Renderer.Copy(s.Quit.texture, &s.Quit.size, &s.Quit.pos)

	Renderer.Present()
}
