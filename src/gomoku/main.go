package main

import (
	"gomoku/gui"
)

func main() {
	gui.StartupGUI()
	defer gui.ShutdownGUI()

	gui.Run()
	return
}
