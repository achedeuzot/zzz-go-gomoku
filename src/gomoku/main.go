package main

import (
	"gomoku/gui"
	"runtime"
)

func init() {
	runtime.LockOSThread()
}

func main() {
	gui.StartupGUI()
	defer gui.ShutdownGUI()

	gui.Run()
	return
}
