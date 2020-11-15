package domain

import (
	"net/http"
	"testing"
)

func TestGetUserNoUserFound(t *testing.T) {
	user, err := GetUser(0)

	if user != nil {
		t.Error("A user was not expected with ID 0")
	}
	if err == nil {
		t.Error("An error was expect for user ID 0")
	}
	if err.StatusCode != http.StatusNotFound {
		t.Error("A StatusNotFound 404 status was expected")
	}
}
