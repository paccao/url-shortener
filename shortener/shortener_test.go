package shortener_test

import (
	"testing"

	"github.com/paccao/url-shortener/shortener"
)

func TestValidateUserInput(t *testing.T) {
	s := "htt://user:pass@host.com:5432/path?k=v#f"

	result := shortener.ValidateUserInput(s)

	if result == s {
		t.Fatal("Url was invalid")
	}
}
