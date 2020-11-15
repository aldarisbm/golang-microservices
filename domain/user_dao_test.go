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
