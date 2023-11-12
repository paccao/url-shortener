package db

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
)

type URL struct {
	URL string `json:"url"`
}

func GetUrl(url string) {
	fileName := "db/urls.json"
	jsonFile, jsonReadErr := os.Open(fileName)
	if jsonReadErr != nil {
		fmt.Println("Error opening file: ", jsonReadErr)
		os.Exit(0)
	}
	defer jsonFile.Close()
	fmt.Printf("Successfully read %s\n", fileName)

	urls, decodeErr := decodeJsonFile(jsonFile)
	if decodeErr != nil {
		fmt.Println("Unexpected error while decoding json file")
		os.Exit(0)
	}

	// TODO: find url in urls

}

func decodeJsonFile(file *os.File) ([]URL, error) {
	var urls []URL
	decoder := json.NewDecoder(file)
	decodeJsonErr := decoder.Decode(&urls)
	if decodeJsonErr != nil {
		fmt.Println("Error decoding JSON: ", decodeJsonErr)
		return nil, errors.New("error decoding JSON from file")
	}
	return urls, nil
}
