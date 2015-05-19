package gui

import (
	"log"

	"github.com/veandco/go-sdl2/sdl"
	"github.com/veandco/go-sdl2/sdl_ttf"
)

func GetTextureFromFont(fontname string, text string, size int, color sdl.Color) *Text {
	ttf.Init()

	font, err := ttf.OpenFont(fontname, size)
	if err != nil {
		log.Fatalf("Failed to load font: %s\n", err)
	}

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

	return &Text{
		texture: text_texture,
		size:    text_surf_size,
	}
}
