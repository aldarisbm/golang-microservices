package services

import "github.com/aldarisbm/golang-microservices/domain"

func GetUser(userID int64) (*domain.User, error) {
	return domain.GetUser(userID)
}
