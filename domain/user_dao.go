package domain

import (
	"fmt"
)

var (
	users = map[int64]*User{
		123: {id: 1, FirstName: "Hello", LastName: "World", Email: "hello@world.com"},
	}
)

func GetUser(userID int64) (*User, error) {
	if user := users[userID]; users != nil {
		return user, nil
	}
	return nil, fmt.Errorf("user %v was not found", userID)
}
