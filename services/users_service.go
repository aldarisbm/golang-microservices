package services

import (
	"github.com/aldarisbm/golang-microservices/domain"
	"github.com/aldarisbm/golang-microservices/utils"
)

type usersService struct{}

var (
	UsersService usersService
)

func (u *usersService) GetUser(userID int64) (*domain.User, *utils.ApplicationError) {
	user, err := domain.UserDao.GetUser(userID)
	if err != nil {
		return nil, err
	}
	return user, nil
}
