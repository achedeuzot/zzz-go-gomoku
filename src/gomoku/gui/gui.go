package gui

import (
	"log"

	"github.com/veandco/go-sdl2/sdl"
)

var (
	winTitle    string = "Go-Gomoku"
	Window      *sdl.Window
	DisplayMode *sdl.DisplayMode
	Renderer    *sdl.Renderer
	Running     bool = true
)

func StartupGUI() {
	err := sdl.VideoInit("")
	if err != nil {
		log.Fatalf("Failed to init video: %s\n", err)
	}

	var dispMode sdl.DisplayMode
	err = sdl.GetDesktopDisplayMode(0, &dispMode)
	if err != nil {
		log.Fatalf("Failed to get desktop display mode: %s\n", err)
	}

	Window, err := sdl.CreateWindow(winTitle, sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED,
		int(dispMode.W), int(dispMode.H), sdl.WINDOW_SHOWN)
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

func isMouseButtonLeftUp(t *sdl.MouseButtonEvent) bool {
	if t.Type == sdl.MOUSEBUTTONUP && t.Button == sdl.BUTTON_LEFT {
		return true
	}
	return false
}
