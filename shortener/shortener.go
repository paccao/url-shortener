package shortener

import (
	"fmt"
	"log"
)

func ShortenUrl() string {
	fmt.Println("input text")
	var w1 string
	n, err := fmt.Scanln(&w1)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("number of items read: %d\n", n)
	fmt.Printf("read line: %s-\n", w1)
	return "w1"
}
