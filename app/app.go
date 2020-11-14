package app

import (
	"log"
	"net/http"

	"github.com/aldarisbm/golang-microservices/controllers"
)

func StartApp() {
	http.HandleFunc("/users", controllers.GetUser)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
