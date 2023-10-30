package shortener_test

import (
	"testing"

	"github.com/paccao/url-shortener/shortener"
)

func TestShortenUrl(t *testing.T) {
	if shortener.ShortenUrl() != "Go Gopher" {
		t.Fatal("Wrong mascot :(")
	}
}
