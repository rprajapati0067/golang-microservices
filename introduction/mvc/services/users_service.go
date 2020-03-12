package services

import (
	"github.com/rprajapati0067/golang-microservices/introduction/mvc/domain"
	"github.com/rprajapati0067/golang-microservices/introduction/mvc/utils"
)

func GetUser(userId int64) (*domain.User, *utils.ApplicationError) {
	return domain.GetUser(userId)
}
