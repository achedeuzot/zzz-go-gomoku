package main

import (
	"gomoku/gui"
)

func main() {
	gui.StartUpGUI()
	defer gui.ShutdownGUI()

	gui.Run()
	return
}
