package gui

import (
	"gomoku/ai"
	"gomoku/arena"
	"gomoku/human"

	"github.com/veandco/go-sdl2/sdl"
	"github.com/veandco/go-sdl2/sdl_ttf"
)

type Options struct {
	Background   *Texture
	Font         *ttf.Font
	Title        *Texture
	AIvsHuman    *Texture
	HumanvsHuman *Texture
	AIvsAI       *Texture
}

func NewOptions() *Options {
	menu := &Options{
		Background:   GetTextureFromImage("data/img/bg.jpg"),
		Title:        GetTextureFromFont("data/fonts/TaiLeb.ttf", "Options", 150, sdl.Color{R: 255, G: 255, B: 255, A: 255}),
		AIvsHuman:    GetTextureFromFont("data/fonts/TaiLeb.ttf", "AI vs Human", 150, sdl.Color{R: 255, G: 255, B: 255, A: 255}),
		HumanvsHuman: GetTextureFromFont("data/fonts/TaiLeb.ttf", "Human vs Human", 150, sdl.Color{R: 255, G: 255, B: 255, A: 255}),
		AIvsAI:       GetTextureFromFont("data/fonts/TaiLeb.ttf", "AI vs AI", 150, sdl.Color{R: 255, G: 255, B: 255, A: 255}),
	}

	menu.Background.pos = sdl.Rect{X: 0, Y: 0, W: DisplayMode.W, H: DisplayMode.H}
	menu.Title.pos = sdl.Rect{X: DisplayMode.W/2 - (menu.Title.size.W*DisplayMode.W/2560)/2, Y: DisplayMode.H / 7, W: menu.Title.size.W * DisplayMode.W / 2560, H: menu.Title.size.H * DisplayMode.H / 1440}
	menu.AIvsHuman.pos = sdl.Rect{X: DisplayMode.W/2 - (menu.AIvsHuman.size.W*DisplayMode.W/2560)/2, Y: (DisplayMode.H / 7) * 3, W: menu.AIvsHuman.size.W * DisplayMode.W / 2560, H: menu.AIvsHuman.size.H * DisplayMode.H / 1440}
	menu.HumanvsHuman.pos = sdl.Rect{X: DisplayMode.W/2 - (menu.HumanvsHuman.size.W*DisplayMode.W/2560)/2, Y: (DisplayMode.H / 7) * 4, W: menu.HumanvsHuman.size.W * DisplayMode.W / 2560, H: menu.HumanvsHuman.size.H * DisplayMode.H / 1440}
	menu.AIvsAI.pos = sdl.Rect{X: DisplayMode.W/2 - (menu.AIvsAI.size.W*DisplayMode.W/2560)/2, Y: (DisplayMode.H / 7) * 5, W: menu.AIvsAI.size.W * DisplayMode.W / 2560, H: menu.AIvsAI.size.H * DisplayMode.H / 1440}
	return menu
}

func (s *Options) handleEvents() {
	for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
		switch t := event.(type) {
		case *sdl.QuitEvent:
			Running = false
		case *sdl.KeyUpEvent:
			if t.Keysym.Sym == sdl.K_ESCAPE {
				CurrScene = SceneMap["MenuMain"]
			}
		case *sdl.MouseButtonEvent:
			if isMouseButtonLeftUp(t) {
				if XYInRect(s.AIvsHuman.pos, t.X, t.Y) {
					player1 := human.NewHuman(arena.BlackPlayer)
					player2 := ai.NewAI(arena.WhitePlayer)
					arena.Gomoku = arena.NewArena(player1, player2)
					CurrScene = SceneMap["Game"]
					break
				} else if XYInRect(s.HumanvsHuman.pos, t.X, t.Y) {
					player1 := human.NewHuman(arena.BlackPlayer)
					player2 := human.NewHuman(arena.WhitePlayer)
					arena.Gomoku = arena.NewArena(player1, player2)
					CurrScene = SceneMap["Game"]
					break
				} else if XYInRect(s.AIvsAI.pos, t.X, t.Y) {
					player1 := ai.NewAI(arena.BlackPlayer)
					player2 := ai.NewAI(arena.WhitePlayer)
					arena.Gomoku = arena.NewArena(player1, player2)
					CurrScene = SceneMap["Game"]
					break
				}
			}
		}
	}
}

func (s *Options) PlayScene() {
	s.handleEvents()

	Renderer.Clear()
	Renderer.SetDrawColor(0, 0, 0, 255)
	Renderer.FillRect(&sdl.Rect{X: 0, Y: 0, W: DisplayMode.W, H: DisplayMode.H})

	Renderer.Copy(s.Background.texture, &s.Background.size, &s.Background.pos)
	Renderer.Copy(s.Title.texture, &s.Title.size, &s.Title.pos)
	Renderer.Copy(s.AIvsHuman.texture, &s.AIvsHuman.size, &s.AIvsHuman.pos)
	Renderer.Copy(s.HumanvsHuman.texture, &s.HumanvsHuman.size, &s.HumanvsHuman.pos)
	Renderer.Copy(s.AIvsAI.texture, &s.AIvsAI.size, &s.AIvsAI.pos)

	Renderer.Present()
}
