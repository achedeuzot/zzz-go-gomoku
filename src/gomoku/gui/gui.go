package gui

import (
	"fmt"
	"os"

	"github.com/veandco/go-sdl2/sdl"
)

var (
	winTitle            string = "Go-Gomoku"
	winWidth, winHeight int    = 2560, 1440
	Window              *sdl.Window
	DisplayMode         *sdl.DisplayMode
	Renderer            *sdl.Renderer
	Running             bool = true
)

func StartupGUI() {
	Window, err := sdl.CreateWindow(winTitle, sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED,
		winWidth, winHeight, sdl.WINDOW_FULLSCREEN)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to create window: %s\n", err)
		os.Exit(1)
	}

	// Using := here seems to be very problematic...
	Renderer, err = sdl.CreateRenderer(Window, -1, sdl.RENDERER_ACCELERATED)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to create renderer: %s\n", err)
		os.Exit(2)
	}

	dispIdx, err := Window.GetDisplayIndex()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to get display mode: %s\n", err)
		os.Exit(2)
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
