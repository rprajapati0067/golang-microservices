package app

import (
	"github.com/rprajapati0067/golang-microservices/introduction/src/api/controllers/polo"
	"github.com/rprajapati0067/golang-microservices/introduction/src/api/controllers/repositories"
)

func mapUrls() {

	router.GET("marco", polo.Polo)
	router.POST("/repository", repositories.CreateRepo)
	router.POST("/repositories", repositories.CreateRepos)

}
