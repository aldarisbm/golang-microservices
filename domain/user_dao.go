package domain

import (
	"fmt"
	"net/http"

	"github.com/aldarisbm/golang-microservices/utils"
)

var (
	users = map[int64]*User{
		123: {ID: 123, FirstName: "Hello", LastName: "World", Email: "hello@world.com"},
	}
	UserDao userDaoInterface
)

func init() {
	UserDao = &userDao{}
}

// we are making an interface so that its easier to mock the underlying functionality
type userDaoInterface interface {
	GetUser(int64) (*User, *utils.ApplicationError)
}

type userDao struct{}

func (u *userDao) GetUser(userID int64) (*User, *utils.ApplicationError) {
	if user := users[userID]; user != nil {
		return user, nil
	}
	return nil, &utils.ApplicationError{
		Message:    fmt.Sprintf("user %v does not exist", userID),
		StatusCode: http.StatusNotFound,
		Code:       "not_found",
	}
}
