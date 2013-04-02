package main

import (
	"bufio"
	"code.google.com/p/freetype-go/freetype"
	"code.google.com/p/freetype-go/freetype/truetype"
	"flag"
	"image"
	"image/draw"
	"image/png"
	"io/ioutil"
	"os"
)

var (
	outputFile = flag.String("output", "out.png", "path to ouput file")
	key        = flag.String("key", "", "API key")
)

type Context struct {
	Img         draw.Image
	FontContext *freetype.Context
	Font        *truetype.Font
	FontSize    float64
}

func Render(levels []Level) {
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
	for i, levelBox := range levelBoxes {
		levelBox.Render(&context, i*120, 0, 100)
	}

	file, _ := os.Create(*outputFile)
	defer file.Close()

	b := bufio.NewWriter(file)
	png.Encode(b, context.Img)
	b.Flush()
}

func main() {
	flag.Parse()
	database := CreateDatabaseFor(*key, 2, 3)
	levels := database.GetLevels()
	Render(levels)
}
