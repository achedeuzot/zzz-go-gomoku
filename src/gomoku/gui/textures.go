package gui

import (
	"log"

	"github.com/veandco/go-sdl2/sdl"
	img "github.com/veandco/go-sdl2/sdl_image"
)

type Texture struct {
	texture *sdl.Texture
	size    sdl.Rect
	pos     sdl.Rect
}

func GetTextureFromImage(filename string) *Texture {
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

	return &Texture{
		texture: texture,
		size:    text_surf_size,
	}
}

func GetTextureFromFont(fontId int8, text string, size int, color sdl.Color) *Texture {
	font := Fonts[fontId]

	text_rendered, err := font.RenderUTF8_Blended(text, color)
	if err != nil {
		log.Fatalf("Failed to render text: %s\n", err)
	}

	var text_surf_size sdl.Rect
	text_rendered.GetClipRect(&text_surf_size)

	text_texture, err := Renderer.CreateTextureFromSurface(text_rendered)
	if err != nil {
		log.Fatalf("Failed to create texture: %s\n", err)
	}

	return &Texture{
		texture: text_texture,
		size:    text_surf_size,
	}
}
