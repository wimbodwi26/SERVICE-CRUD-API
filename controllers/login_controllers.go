package controllers

import (
	"backend-go/database"
	"backend-go/helpers"
	"backend-go/models"
	"backend-go/structs"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func Login(c *gin.Context) {
	var req = structs.UserLoginRequest{}
	var user = models.User{}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusUnprocessableEntity, structs.ErrorResponse{
			Success: false,
			Message: "validation errors",
			Errors: helpers.TranslateErrorMessage(err),
		})
		return
	}

	if err := database.DB.Where("username= ?", req.Username).First(&user).Error; err != nil {
		c.JSON(http.StatusUnauthorized, structs.ErrorResponse{
			Success: false,
			Message: "User Not Found",
			Errors: helpers.TranslateErrorMessage(err),
		})
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
		c.JSON(http.StatusUnauthorized, structs.ErrorResponse{
			Success: false,
			Message: "Invalid Password",
			Errors: helpers.TranslateErrorMessage(err),
		})
		return
	}

	token := helpers.GenerateToken(user.Username)

	c.JSON(http.StatusOK, structs.SuccessResponse{
		Succes: 	true,
		Message: 	"Login Success",
		Data: structs.UserResponse{
			Id: 		user.Id,
			Name: 		user.Name,
			Username: 	user.Username,
			Email: 		user.Email,
			CreatedAt: 	user.CreatedAt.String(),
			UpdatedAt: 	user.UpdatedAt.String(),
			Token: 		&token,
		}, 
	})
}