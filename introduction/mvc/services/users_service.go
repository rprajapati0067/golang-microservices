package services

import (
	"github.com/rprajapati0067/golang-microservices/introduction/mvc/domain"
	"github.com/rprajapati0067/golang-microservices/introduction/mvc/utils"
	"log"
)

type userService struct {
}

var (
	UserService userService
)

func (u *userService) GetUser(userId int64) (*domain.User, *utils.ApplicationError) {
	log.Println("Service called")
	user, err := domain.UserDao.GetUser(userId)
	if err != nil {
		return nil, err
	}
	return user, nil
}
 