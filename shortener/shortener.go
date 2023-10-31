package shortener

import (
	"fmt"
	"log"
	"net/url"
	"regexp"
)

func ParseUserInput() string {
	var w1 string
	n, err := fmt.Scanf("%f\n", &w1)
	if err != nil || n != 1 {
		log.Fatal("No input was passed", err)
	}

	return w1
}

func ValidateUserInput(userInputUrl string) string {
	u, parseErr := url.Parse(userInputUrl)
	if parseErr != nil {
		log.Fatal(parseErr)
	}
	if u.Scheme == "" || u.Host == "" || u.Path == "" {
		log.Fatalf("Error! Either the scheme, the host or the path was empty. Scheme: '%s', Host: '%s', Path: '%s'", u.Scheme, u.Host, u.Path)
	}
	schemeMatch, regErr := regexp.MatchString("https|http", u.Scheme)
	if regErr != nil {
		log.Fatal(regErr)
	}

	if !schemeMatch {
		log.Fatalf("Url scheme didnt match https or http. Scheme: '%s'", u.Scheme)
	}

	log.Println("Success:", schemeMatch, u)
	return userInputUrl
}
