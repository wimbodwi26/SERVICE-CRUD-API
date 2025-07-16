package controllers

import (
	"net/http"
	"backend-go/database"
	"backend-go/helpers"
	"backend-go/models"
	"backend-go/structs"

	"github.com/gin-gonic/gin"
)

func Register(c *gin.Context) {
	var req = structs.UserCreateRequest{}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusUnprocessableEntity, structs.ErrorResponse{
			Success: false,
			Message: "validasi Errors",
			Errors: helpers.TranslateErrorMessage(err),
		})
		return
	}

	user := models.User{
		Name: 		req.Name,
		Username: 	req.Username,
		Email: 		req.Email,
		Password: 	helpers.HashPassword(req.Password),
	}

	if err := database.DB.Create(&user).Error; err != nil {
		if helpers.IsDuplicateEntryError(err) {
			c.JSON(http.StatusConflict, structs.ErrorResponse{
				Success: 	false,
				Message: 	"Duplicate entry error",
				Errors: 	helpers.TranslateErrorMessage(err),
			})
		} else {
			c.JSON(http.StatusInternalServerError, structs.ErrorResponse{
				Success: 	false,
				Message: 	"Failed to create user",
				Errors: 	helpers.TranslateErrorMessage(err),
			})
		}
		return
	}

	c.JSON(http.StatusCreated, structs.SuccessResponse{
		Succes: 	true,
		Message: 	"User created successfully",
		Data:  		structs.UserResponse{
			Id:			user.Id,
			Name: 		user.Name,
			Username: 	user.Username,
			Email: 		user.Email,
			CreatedAt:  user.CreatedAt.Format("2006-01-02 15:04:05"),
			UpdatedAt:  user.UpdatedAt.Format("2006-01-02 15:04:05"),
		}, 
	})
}