package main

import (
	"gomoku/gui"
)

func main() {
	gui.StartUpGUI()
	defer gui.ShutdownGUI()
	return
}
