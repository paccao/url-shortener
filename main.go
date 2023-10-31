package main

import (
	"fmt"

	"github.com/paccao/url-shortener/shortener"
)

func main() {
	fmt.Println("Please input your url here: ")

	userInput := shortener.ParseUserInput()

	shortener.ValidateUserInput(userInput)

}
