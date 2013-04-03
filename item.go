package main

import (
	"image/color"
)

const (
	statusApprentice  = 1
	statusGuru        = 2
	statusMaster      = 3
	statusEnlightened = 4
	statusBurned      = 5
	statusUnknown     = 6
)

const (
	typeRadical    = 1
	typeKanji      = 2
	typeVocabulary = 3
	typeUnknown    = 4
)

type ItemStatus int
type ItemType int

func GetStatusFromString(stringStatus string) ItemStatus {
	switch stringStatus {
	case "apprentice":
		return statusApprentice
	case "guru":
		return statusGuru
	case "master":
		return statusMaster
	case "enlighten":
		return statusEnlightened
	case "burned":
		return statusBurned
	default:
		return statusUnknown
	}

	return statusUnknown
}

func GetStringFromStatus(status ItemStatus) string {
	switch status {
	case statusApprentice:
		return "appr"
	case statusGuru:
		return "guru"
	case statusMaster:
		return "mstr"
	case statusEnlightened:
		return "enli"
	case statusBurned:
		return "brnd"
	default:
		return "unkn"
	}

	return "unknown"
}

func GetStringForItemType(itemType ItemType) string {
	switch itemType {
	case typeRadical:
		return "radicals"
	case typeKanji:
		return "kanji"
	case typeVocabulary:
		return "vocabulary"
	case typeUnknown:
		return "unknown"
	default:
		panic("Unknown ItemType")
	}

	return "unknown"
}

type Item struct {
	Status ItemStatus
	Level  int
	Name   string
	Type   ItemType
}

func NewItemFromJson(json RequestedInformation) Item {
	character := json.Character

	if character == "" {
		character = "nil"
	}

	stringStatus := json.Stats.Srs
	status := GetStatusFromString(stringStatus)
	level := int(json.Level)

	return Item{status, level, character, typeUnknown}
}

func (item *Item) Color() color.RGBA {
	switch item.Status {
	case statusUnknown:
		return color.RGBA{40, 40, 40, 255}
	case statusApprentice:
		return color.RGBA{221, 0, 147, 255}
	case statusGuru:
		return color.RGBA{136, 45, 158, 255}
	case statusMaster:
		return color.RGBA{41, 77, 219, 255}
	case statusEnlightened:
		return color.RGBA{0, 147, 221, 255}
	case statusBurned:
		return color.RGBA{74, 74, 74, 255}
	default:
		return color.RGBA{255, 0, 0, 255}
	}

	return color.RGBA{0, 0, 0, 0}
}

func (item *Item) TypeColor() color.RGBA {
	switch item.Type {
	case typeRadical:
		return color.RGBA{0, 147, 221, 255}
	case typeKanji:
		return color.RGBA{221, 0, 147, 255}
	case typeVocabulary:
		return color.RGBA{136, 45, 158, 255}
	default:
		return color.RGBA{255, 0, 0, 255}
	}

	return color.RGBA{0, 0, 0, 0}
}
