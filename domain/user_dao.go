package domain

import (
	"fmt"
	"net/http"

	"github.com/antonimassomola/golang-microservices/mvc/utils"
)

var (
	users = map[uint64]*User{
		123: &User{Id: 123, FirstName: "Alex", LastName: "Massó", Email: "alex@gmail.com"},
	}
)

// GetUser to retrieve from DB
func GetUser(userID uint64) (*User, *utils.ApplicationError) {
	if user := users[userID]; user != nil {
		return user, nil
	}

	return nil, &utils.ApplicationError{
		Message:    fmt.Sprintf("User %v was not found", userID),
		StatusCode: http.StatusNotFound,
		Code:       "not_found",
	}
}
