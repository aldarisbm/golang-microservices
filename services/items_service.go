package services

import (
	"net/http"

	"github.com/aldarisbm/golang-microservices/domain"
	"github.com/aldarisbm/golang-microservices/utils"
)

type itemsService struct{}

var (
	ItemsService itemsService
)

func (i *itemsService) GetItem(itemID string) (*domain.Item, *utils.ApplicationError) {
	return nil, &utils.ApplicationError{
		Message:    "implement me",
		StatusCode: http.StatusInternalServerError,
		Code:       "not_implemented",
	}
}
