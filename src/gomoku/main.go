package main

import (
	"flag"
	"gomoku/gui"
	"runtime"

	"github.com/davecheney/profile"
)

var (
	flag_fullscreen bool
	flag_width      int
	flag_height     int
)

func init() {
	runtime.LockOSThread()

	flag.BoolVar(&flag_fullscreen, "f", true, "Enable fullscreen")
	flag.IntVar(&flag_width, "w", 0, "Width")
	flag.IntVar(&flag_height, "h", 0, "Height")
	flag.Parse()
}

func main() {
	defer profile.Start(profile.CPUProfile).Stop()
	gui.StartupGUI(flag_fullscreen, flag_width, flag_height)
	defer gui.ShutdownGUI()

	gui.Run()
	return
}
