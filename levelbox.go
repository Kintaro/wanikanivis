package main

type LevelBox struct {
	Boxes []StatusBox
}

func NewLevelBox(level *Level) *LevelBox {
	result := &LevelBox{make([]StatusBox, 5)}
	for i := range result.Boxes {
		result.Boxes[i] = CreateFromLevel(level, ItemStatus(i+1))
	}
	return result
}

func (levelBox *LevelBox) Render(context *Context, x int, y int, w int) {
	for _, box := range levelBox.Boxes {
		if len(box.Boxes) == 0 {
			continue
		}
		box.Render(context, x, y, w)
		y += box.GetHeight()
		y += box.HeaderBox.GetHeight()
	}
}

func (levelBox *LevelBox) GetHeight() int {
	sum := 0
	for _, box := range levelBox.Boxes {
		sum += box.GetHeight()
	}
	return sum
}
