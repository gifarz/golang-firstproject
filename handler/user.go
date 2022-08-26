package handler

import (
	"first-project/helper"
	"first-project/user"
	"net/http"

	"github.com/gin-gonic/gin"
)

type userHandler struct {
	userService user.Service
}

func NewUserHandler(userService user.Service) *userHandler {
	return &userHandler{userService}
}

func (uh *userHandler) RegisterUser(c *gin.Context) {
	var input user.RegisterUserInput

	err := c.ShouldBindJSON(&input)
	if err != nil {
		formatError := helper.FormatError(err)

		dataFormatter := gin.H{"error": formatError}

		apiResponse := helper.ResponseAPI(
			"Registration failed!",
			http.StatusUnprocessableEntity,
			"error",
			dataFormatter,
		)
		c.JSON(http.StatusUnprocessableEntity, apiResponse)
		return
	} else {
		userRes, err := uh.userService.RegisterUser(input)

		if err != nil {

			apiResponse := helper.ResponseAPI(
				"Registration failed!",
				http.StatusBadRequest,
				"error",
				err,
			)
			c.JSON(http.StatusBadRequest, apiResponse)
			return
		} else {
			dataFormatter := user.Formatter(userRes, "tokenjwt")
			apiResponse := helper.ResponseAPI(
				"User has been registered",
				http.StatusOK,
				"success",
				dataFormatter,
			)
			c.JSON(http.StatusOK, apiResponse)
			return
		}
	}

}
