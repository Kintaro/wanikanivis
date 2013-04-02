package main

import (
	"container/list"
)

type Level struct {
	Items *list.List
}

func NewLevel() Level {
	return Level{list.New()}
}

func (level *Level) FilterForType(itemType ItemType) Level {
	newLevel := NewLevel()

	for i := level.Items.Front(); i != nil; i = i.Next() {
		item := i.Value.(*Item)
		if item.Type == itemType {
			newLevel.Items.PushBack(item)
		}
	}

	return newLevel
}

func (level *Level) FilterForStatus(itemStatus ItemStatus) Level {
	newLevel := NewLevel()

	for i := level.Items.Front(); i != nil; i = i.Next() {
		item := i.Value.(*Item)
		if item.Status == itemStatus {
			newLevel.Items.PushBack(item)
		}
	}

	return newLevel
}
