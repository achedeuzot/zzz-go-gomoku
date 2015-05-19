package gui

import (
	"log"

	"github.com/veandco/go-sdl2/sdl"
	"github.com/veandco/go-sdl2/sdl_image"
)

func loadPng(filename string) *sdl.Texture {
	var surface *sdl.Surface
	var texture *sdl.Texture
	var err error

	img.Init(img.INIT_PNG)

	surface, err = img.Load(filename)
	if err != nil {
		log.Fatalf("Failed to load PNG: %s\n", err)
	}
	defer surface.Free()

	texture, err = Renderer.CreateTextureFromSurface(surface)
	if err != nil {
		log.Fatalf("Failed to create texture: %s\n", err)
	}

	return texture
}
