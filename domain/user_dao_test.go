package domain

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetUserNoUserFound(t *testing.T) {
	user, err := GetUser(0)

	assert.Nil(t, user, "user was not expected")
	assert.NotNil(t, err, "were expecting an error but didnt get one")
	assert.EqualValues(t, err.StatusCode, http.StatusNotFound)
	assert.EqualValues(t, "not_found", err.Code)
	assert.EqualValues(t, "user 0 does not exist", err.Message)
}

func TestGetUserUserFound(t *testing.T) {
	user, err := GetUser(123)

	assert.Nil(t, err, "error was not expected")
	assert.NotNil(t, user, "User was expected, and not found")

	assert.EqualValues(t, 123, user.ID)
	assert.EqualValues(t, "Hello", user.FirstName)
	assert.EqualValues(t, "World", user.LastName)
	assert.EqualValues(t, "hello@world.com", user.Email)
}
