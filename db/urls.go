package db

import (
	"encoding/json"
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
}
