package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/rprajapati0067/golang-microservices/introduction/mvc/services"
	"github.com/rprajapati0067/golang-microservices/introduction/mvc/utils"
	"net/http"
	"strconv"
)

func GetUser(c *gin.Context) {

	userId, err := strconv.ParseInt(c.Param("user_id"), 10, 64)
	if err != nil {

		apiErr := &utils.ApplicationError{
			Message:    "user_id must be a number",
			StatusCode: http.StatusBadRequest,
			Code:       "bad_request",
		}
		utils.Respond(c, apiErr.StatusCode, apiErr)
		return
	}

	user, apiErr := services.UserService.GetUser(userId)

	if apiErr != nil {
		utils.Respond(c, apiErr.StatusCode, apiErr)
		return
	}
	utils.Respond(c, http.StatusOK, user)

}
