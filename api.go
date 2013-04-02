package main

import (
	"io/ioutil"
	"log"
	"net/http"
)

func GetDataFromApi(key string, apitype string) []byte {
	url := "http://www.wanikani.com/api/user/" + key + "/" + apitype

	response, err := http.Get(url)

	if err != nil {
		log.Fatalln("Error: http.Get.", err)
	}

	content, err := ioutil.ReadAll(response.Body)
	defer response.Body.Close()

	if err != nil {
		log.Fatalln("Error: ioutil.ReadAll.", err)
	}

	return content
}

func CreateDatabaseFor(key string, args ...ItemType) Database {
	database := NewDatabase()

	for _, v := range args {
		tempDatabase := NewDatabase()
		data := GetDataFromApi(key, GetStringForItemType(v))
		tempDatabase.LoadFromData(data, v)
		database.Join(tempDatabase)
	}

	return database
}
