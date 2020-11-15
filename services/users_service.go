package services

import (
	"github.com/aldarisbm/golang-microservices/domain"
	"github.com/aldarisbm/golang-microservices/utils"
)

func GetUser(userID int64) (*domain.User, *utils.ApplicationError) {
	return domain.GetUser(userID)
}
