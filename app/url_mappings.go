package app

import (
	"github.com/aldarisbm/golang-microservices/controllers"
)

// one place to look at all url mappings
func mapUrls() {
	router.GET("/users/:user_id", controllers.GetUser)
}
