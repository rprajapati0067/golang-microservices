package repositories

import (
	"github.com/gin-gonic/gin"
	"github.com/rprajapati0067/golang-microservices/introduction/src/api/domain/repositories"
	"github.com/rprajapati0067/golang-microservices/introduction/src/api/services"
	"github.com/rprajapati0067/golang-microservices/introduction/src/api/utils/errors"
	"net/http"
)

func CreateRepo(c *gin.Context) {
	var request repositories.CreateRepoRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		apiErr := errors.NewBadRequestError("invalid json body")
		c.JSON(apiErr.Status(), apiErr)
		return
	}

	result, err := services.RepositoryService.CreateRepo(request)
	if err != nil {
		c.JSON(err.Status(), err)
		return
	}
	c.JSON(http.StatusCreated, result)

}
func CreateRepos(c *gin.Context) {
	var request []repositories.CreateRepoRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		apiErr := errors.NewBadRequestError("invalid json body")
		c.JSON(apiErr.Status(), apiErr)
		return
	}

	result, err := services.RepositoryService.CreateRepos(request)
	if err != nil {
		c.JSON(err.Status(), err)
		return
	}
	c.JSON(result.StatusCode, result)

}
