package gui

import (
	"log"

	"github.com/veandco/go-sdl2/sdl"
	"github.com/veandco/go-sdl2/sdl_image"
)

type Text struct {
	texture *sdl.Texture
	size    sdl.Rect
}

func GetTextureFromImage(filename string) *Text {
	var surface *sdl.Surface
	var texture *sdl.Texture
	var err error

	surface, err = img.Load(filename)
	if err != nil {
		log.Fatalf("Failed to load PNG: %s\n", err)
	}
	defer surface.Free()

	var text_surf_size sdl.Rect
	surface.GetClipRect(&text_surf_size)

	texture, err = Renderer.CreateTextureFromSurface(surface)
	if err != nil {
		log.Fatalf("Failed to create texture: %s\n", err)
	}

	return &Text{
		texture: texture,
		size:    text_surf_size,
	}
}
