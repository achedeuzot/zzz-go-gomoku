package gui

import (
	"fmt"
	"os"

	"github.com/veandco/go-sdl2/sdl"
)

var (
	winTitle            string = "Go-Gomoku"
	winWidth, winHeight int    = 800, 600
	Window              *sdl.Window
	Renderer            *sdl.Renderer
	Running             bool = true
)

func StartUpGUI() *sdl.Renderer {
	Window, err := sdl.CreateWindow(winTitle, sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED,
		winWidth, winHeight, sdl.WINDOW_SHOWN)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to create window: %s\n", err)
		os.Exit(1)
	}

	Renderer, err := sdl.CreateRenderer(Window, -1, sdl.RENDERER_ACCELERATED)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to create renderer: %s\n", err)
		os.Exit(2)
	}

	return Renderer
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
	for Running {
		(*CurrScene).PlayScene()
	}
}
