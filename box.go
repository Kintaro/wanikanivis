package main

import (
	"code.google.com/p/freetype-go/freetype"
	"image"
	"image/color"
	"image/draw"
)

type Box struct {
	Content   string
	Color     color.RGBA
	TypeColor color.RGBA
}

func (box *Box) Render(context *Context, x int, y int, w int) {
	rectangle := image.Rect(x, y, x+w, y+box.GetHeight())
	draw.Draw(context.Img, rectangle, &image.Uniform{&box.Color}, image.ZP, draw.Src)
	typeRectangle := image.Rect(x, y, x+10, y+box.GetHeight())
	draw.Draw(context.Img, typeRectangle, &image.Uniform{&box.TypeColor}, image.ZP, draw.Src)
	box.RenderText(context, x+10, y)
}

func (box *Box) RenderText(context *Context, x int, y int) {
	pt := freetype.Pt(x, y)
	pt.Y += context.FontContext.PointToFix32(18.0)
	context.FontContext.SetSrc(&image.Uniform{color.RGBA{255, 255, 255, 255}})

	for _, s := range box.Content {
		context.FontContext.DrawString(string(s), pt)
		pt.X += context.FontContext.PointToFix32(18.0)
	}
	context.FontContext.SetClip(context.Img.Bounds())
}

func (box *Box) GetHeight() int {
	return 20
}
