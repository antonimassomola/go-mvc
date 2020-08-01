package services

import (
	"github.com/antonimassomola/golang-microservices/mvc/domain"
	"github.com/antonimassomola/golang-microservices/mvc/utils"
)

func GetUser(userID int64) (*domain.User, *utils.ApplicationError) {
	return domain.GetUser(userID)
}
