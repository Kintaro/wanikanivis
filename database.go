package main

import (
	"container/list"
	"encoding/json"
)

type Database struct {
	Items *list.List
}

type Statistics struct {
	Srs string
}

type RequestedInformation struct {
	Character string
	Stats     Statistics
	Level     float64
	General   []RequestedInformation
}

type JsonStruct struct {
	Requested_information []RequestedInformation
}

type JsonStructVocabulary struct {
	Requested_information RequestedInformation
}

func NewDatabase() Database {
	return Database{list.New()}
}

func (database *Database) LoadFromData(data []byte, itemType ItemType) {
	jsonData := JsonStruct{}
	err := json.Unmarshal(data, &jsonData)
	if err != nil {
		//panic(err)
	}

	for _, requested_information := range jsonData.Requested_information {
		database.ParseRequestedInformation(requested_information, itemType)
	}

	if len(jsonData.Requested_information) == 0 {
		jsonDataVocabulary := JsonStructVocabulary{}
		err := json.Unmarshal(data, &jsonDataVocabulary)
		if err != nil {
			//panic(err)
		}

		for _, requested_information := range jsonDataVocabulary.Requested_information.General {
			database.ParseRequestedInformation(requested_information, itemType)
		}
	}
}

func (database *Database) ParseRequestedInformation(requested_information RequestedInformation, itemType ItemType) {
	if requested_information.General != nil {
		database.ParseRequestedInformation(requested_information, itemType)
		return
	}
	item := NewItemFromJson(requested_information)
	item.Type = itemType
	database.Items.PushBack(&item)
}

func (database *Database) Join(otherDatabase Database) {
	for i := otherDatabase.Items.Front(); i != nil; i = i.Next() {
		database.Items.PushBack(i.Value)
	}
}

func (database *Database) GetHighestLevel() int {
	max := 0

	for i := database.Items.Front(); i != nil; i = i.Next() {
		item := i.Value.(*Item)
		if item.Level > max {
			max = item.Level
		}
	}

	return max
}

func (database *Database) GetLevels() []Level {
	numberOfLevels := database.GetHighestLevel()
	levels := make([]Level, numberOfLevels)

	for i := 0; i < numberOfLevels; i++ {
		levels[i] = NewLevel()
	}

	for i := database.Items.Front(); i != nil; i = i.Next() {
		item := i.Value.(*Item)
		level := item.Level
		levels[level-1].Items.PushBack(item)
	}

	return levels
}

func (database *Database) GetSize() int {
	return database.Items.Len()
}
