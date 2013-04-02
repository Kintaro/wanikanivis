package main

import (
	"fmt"
	"bufio"
	"code.google.com/p/freetype-go/freetype"
	"code.google.com/p/freetype-go/freetype/truetype"
	"image"
	"image/draw"
	"image/png"
	"io/ioutil"
	"os"
)

type Context struct {
	Img         draw.Image
	FontContext *freetype.Context
	Font        *truetype.Font
	FontSize    float64
}

func main() {
	database := CreateDatabaseFor("123456789", 2, 3)
	fmt.Printf("Database size: %d\n", database.GetSize())
	levels := database.GetLevels()
	levelBoxes := make([]*LevelBox, len(levels))
	maxHeight := 0

	for i, level := range levels {
		levelBoxes[i] = NewLevelBox(&level)
		if levelBoxes[i].GetHeight() > maxHeight {
			maxHeight = levelBoxes[i].GetHeight()
		}
	}

	fontbytes, _ := ioutil.ReadFile("ipag.ttf")
	font, _ := freetype.ParseFont(fontbytes)
	fontcontext := freetype.NewContext()
	fontcontext.SetDPI(72)
	fontcontext.SetFont(font)
	fontcontext.SetFontSize(18)
	img := image.NewRGBA(image.Rect(0, 0, 1920, maxHeight))
	fontcontext.SetDst(img)
	context := Context{img, fontcontext, font, 12.0}

	for i, levelBox := range levelBoxes {
		levelBox.Render(&context, i*120, 0, 100)
	}

	file, _ := os.Create("test.png")
	defer file.Close()

	b := bufio.NewWriter(file)
	png.Encode(b, context.Img)
	b.Flush()
}
