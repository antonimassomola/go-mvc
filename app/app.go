package app

import (
	"fmt"
	"net/http"

	"github.com/antonimassomola/golang-microservices/mvc/controllers"
)

func StartApp() {
	fmt.Println("StartApp")
	http.HandleFunc("/users", controllers.GetUser)
	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}
