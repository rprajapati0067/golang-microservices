package app

import "github.com/rprajapati0067/golang-microservices/introduction/mvc/controllers"

func mapUrls() {

	router.GET("/users/:user_id", controllers.GetUser)
}
