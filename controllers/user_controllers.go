package controllers

import (
	"backend-go/database"
	"backend-go/helpers"
	"backend-go/models"
	"backend-go/structs"
	"net/http"

	"github.com/gin-gonic/gin"
)

func FindUsers(c *gin.Context) {
	var users []models.User

	database.DB.Find(&users)

	c.JSON(http.StatusOK, structs.SuccessResponse{
		Succes: true,
		Message: "List Data Users",
		Data: users,
	})
}

func CraeteUser(c *gin.Context) {
	var req = structs.UserCreateRequest{}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusUnprocessableEntity, structs.ErrorResponse{
			Success: false,
			Message: "Validation Errors",
			Errors: helpers.TranslateErrorMessage(err),
		})
		return
	}

	user := models.User{
		Name: 		req.Name,
		Username: 	req.Username,
		Email: 		req.Email,
		Password: 	req.Password,
	}

	if err := database.DB.Create(&user).Error; err !=nil {
		c.JSON(http.StatusInternalServerError, structs.ErrorResponse{
			Success: false,
			Message: "Faild to create user",
			Errors: helpers.TranslateErrorMessage(err),
		})
		return
	}

	c.JSON(http.StatusCreated, structs.SuccessResponse{
		Succes: true,
		Message: "User Created successfully!",
		Data: structs.UserResponse{
			Id: 		user.Id,
			Name: 		user.Name,
			Username: 	user.Username,
			Email: 		user.Email,
			CreatedAt: 	user.CreatedAt.Format("2006-01-02 15:04:05"),
			UpdatedAt: 	user.UpdatedAt.Format("2006-01-02 15:04:05"),
		},
	})
}

func FindUsersById(c *gin.Context) {
	id := c.Param("id")

	var user models.User

	if err := database.DB.First(&user, id).Error; err != nil {
		c.JSON(http.StatusNotFound, structs.ErrorResponse{
			Success: false,
			Message: "User not found",
			Errors: helpers.TranslateErrorMessage(err),
		})
		return
	}

	c.JSON(http.StatusOK, structs.SuccessResponse{
		Succes: true,
		Message: "User Found",
		Data: structs.UserResponse{
			Id: 	user.Id,
			Name: 	user.Name,
			Email: 	user.Email,
			CreatedAt: 	user.CreatedAt.Format("2006-01-02 15:04:05"),
			UpdatedAt: 	user.UpdatedAt.Format("2006-01-02 15:04:05"), 
		},
	})
}

func UpdateUser(c *gin.Context){
	id := c.Param("id")

	var user models.User

	if err := database.DB.First(&user, id).Error; err != nil {
		c.JSON(http.StatusNotFound, structs.ErrorResponse{
			Success: false,
			Message: "Validation Errors",
			Errors: helpers.TranslateErrorMessage(err),
		})
		return
	}
	
	var req = structs.UserUpdateRequest{}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusUnprocessableEntity, structs.ErrorResponse{
			Success: false,
			Message: "Validation errors",
			Errors: helpers.TranslateErrorMessage(err),
		})
		return
	}

	user.Name = req.Name
	user.Username = req.Username
	user.Email = req.Email
	user.Password = helpers.HashPassword(req.Password)

	if err := database.DB.Save(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, structs.ErrorResponse{
			Success: false,
			Message: "Faild to update user",
			Errors: helpers.TranslateErrorMessage(err),
		})
		return
	}

	c.JSON(http.StatusOK, structs.SuccessResponse{
		Succes: true,
		Message: "User update successfuly",
		Data: structs.UserResponse{
			Id: 		user.Id,
			Name: 		user.Name,
			Username: 	user.Username,
			Email: 		user.Email,
			CreatedAt: 	user.CreatedAt.Format("2006-01-02 15:04:05"),
			UpdatedAt: 	user.UpdatedAt.Format("2006-01-02 15:04:05"),
		},
	})
}

func DeleteUser(c *gin.Context) {
	id := c.Param("id")

	var user models.User

	if err := database.DB.First(&user, id).Error; err != nil {
		c.JSON(http.StatusNotFound, structs.ErrorResponse{
			Success: false,
			Message: "User not found",
			Errors: helpers.TranslateErrorMessage(err),
		})
		return
	}

	if err := database.DB.Delete(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, structs.ErrorResponse{
			Success: false,
			Message: "Faild to delete user",
			Errors: helpers.TranslateErrorMessage(err),
		})
		return
	}

	c.JSON(http.StatusOK, structs.SuccessResponse{
		Succes: true,
		Message: "User delete successfully",
	})
}