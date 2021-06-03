package services

import (
	"github.com/rprajapati0067/golang-microservices/introduction/mvc/domain"
	"github.com/rprajapati0067/golang-microservices/introduction/mvc/utils"
)

type itemService struct {
}

var (
	ItemService itemService
)

func (i *itemService) GetItem(itemId string) (*domain.Item, *utils.ApplicationError) {

	return nil, &utils.ApplicationError{
		Message:    "Not Implemented",
		StatusCode: 404,
		Code:       "no content",
	}
}
