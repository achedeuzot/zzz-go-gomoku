package gui

import (
	"log"

	"github.com/veandco/go-sdl2/sdl"
	ttf "github.com/veandco/go-sdl2/sdl_ttf"
)

var (
	winTitle    string = "Go-Gomoku"
	Window      *sdl.Window
	DisplayMode *sdl.DisplayMode
	Renderer    *sdl.Renderer
	Running     bool = true
	Fonts       map[int8]*ttf.Font
)

func StartupGUI(fullscreen bool, width int, height int) {
	err := sdl.VideoInit("")
	if err != nil {
		log.Fatalf("Failed to init video: %s\n", err)
	}

	var dispMode sdl.DisplayMode
	err = sdl.GetDesktopDisplayMode(0, &dispMode)
	if err != nil {
		log.Fatalf("Failed to get desktop display mode: %s\n", err)
	}

	var sdlflags uint32
	if fullscreen == true {
		sdlflags = sdl.WINDOW_FULLSCREEN_DESKTOP
	} else {
		sdlflags = sdl.WINDOW_SHOWN
	}

	if width > 0 && height > 0 {
		Window, err = sdl.CreateWindow(winTitle, sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED,
			width, height, sdlflags)
	} else {
		Window, err = sdl.CreateWindow(winTitle, sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED,
			int(dispMode.W), int(dispMode.H), sdlflags)
	}
	if err != nil {
		log.Fatalf("Failed to create window: %s\n", err)
	}

	// Using := here seems to be very problematic...
	Renderer, err = sdl.CreateRenderer(Window, -1, sdl.RENDERER_ACCELERATED)
	if err != nil {
		log.Fatalf("Failed to create renderer: %s\n", err)
	}

	dispIdx, err := Window.GetDisplayIndex()
	if err != nil {
		log.Fatalf("Failed to get display mode: %s\n", err)
	}

	DisplayMode = &sdl.DisplayMode{}
	sdl.GetCurrentDisplayMode(dispIdx, DisplayMode)
	if width > 0 && height > 0 && fullscreen == false {
		DisplayMode.W = int32(width)
		DisplayMode.H = int32(height)
	}

	// Setup scaling quality to linear
	sdl.SetHintWithPriority(sdl.HINT_RENDER_SCALE_QUALITY, "1", sdl.HINT_OVERRIDE)

	initFonts()
}

func ShutdownGUI() {
	if Renderer != nil {
		Renderer.Clear()
		Renderer.Destroy()
	}
	if Window != nil {
		Window.Destroy()
	}
}

func Run() {
	initScenes()

	fps := 30
	fpsMill := 1000 / fps

	for Running {
		currentFrameTime := sdl.GetTicks()

		CurrScene.PlayScene()

		currentSpeed := sdl.GetTicks() - currentFrameTime
		if uint32(fpsMill) > currentSpeed {
			sdl.Delay(uint32(fpsMill) - currentSpeed)
		}
	}
}

func XYInRect(rect sdl.Rect, x int32, y int32) bool {
	return ((x > rect.X && x < rect.X+rect.W) && (y > rect.Y && y < rect.Y+rect.H))
}

func isMouseButtonLeftUp(t *sdl.MouseButtonEvent) bool {
	if t.Type == sdl.MOUSEBUTTONUP && t.Button == sdl.BUTTON_LEFT {
		return true
	}
	return false
}

func initFonts() {
	ttf.Init()

	font, err := ttf.OpenFont("data/fonts/TaiLeb.ttf", 70)
	if err != nil {
		log.Fatalf("Failed to load font: %s\n", err)
	}

	Fonts = make(map[int8]*ttf.Font)

	Fonts[0] = font
}
