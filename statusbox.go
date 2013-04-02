package main

import (
	"fmt"
	"image/color"
)

type StatusBox struct {
	Boxes     []*Box
	HeaderBox *Box
}

func CreateFromLevel(level *Level, status ItemStatus) StatusBox {
	filteredLevel := level.FilterForStatus(status)
	var statusBox StatusBox
	statusBox.Boxes = []*Box{}
	statusBox.InitializeWithTypeAndAppend(&filteredLevel, typeRadical)
	statusBox.InitializeWithTypeAndAppend(&filteredLevel, typeKanji)
	statusBox.InitializeWithTypeAndAppend(&filteredLevel, typeVocabulary)
	statusBox.HeaderBox = &Box{GetStringFromStatus(status), color.RGBA{0, 0, 0, 255}, color.RGBA{0, 0, 0, 255}}
	return statusBox
}

func (statusBox *StatusBox) InitializeWithTypeAndAppend(level *Level, itemType ItemType) {
	filteredLevel := level.FilterForType(itemType)
	tempBoxes := make([]*Box, filteredLevel.Items.Len())
	index := 0

	for i := filteredLevel.Items.Front(); i != nil; i = i.Next() {
		item := i.Value.(*Item)
		tempBoxes[index] = &Box{item.Name, item.Color(), item.TypeColor()}
		index++
	}

	statusBox.Boxes = append(statusBox.Boxes, tempBoxes...)
}

func (statusBox *StatusBox) Render(context *Context, x int, y int, w int) {
	statusBox.HeaderBox.Render(context, x, y, w)
	y += statusBox.HeaderBox.GetHeight()
	for _, box := range statusBox.Boxes {
		box.Render(context, x, y, w)
		y += box.GetHeight()
	}
}

func (statusBox *StatusBox) GetHeight() int {
	sum := 0
	for _, box := range statusBox.Boxes {
		sum += box.GetHeight()
	}
	sum += statusBox.HeaderBox.GetHeight()
	return sum
}
