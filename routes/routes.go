package routes

import (
	"backend-go/controllers"
	"backend-go/middlewares"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()

	router.Use(cors.New(cors.Config{
		AllowOrigins:   []string{"*"},
		AllowMethods: 	[]string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders: 	[]string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders: 	[]string{"Content-Length"},
		AllowCredentials: true,
	}))

	router.POST("/api/register", controllers.Register)

	router.POST("/api/login", controllers.Login)

	router.GET("/api/users", middlewares.AuthMiddleware(), controllers.FindUsers)

	router.POST("/api/users", middlewares.AuthMiddleware(), controllers.CraeteUser)

	router.GET("/api/users/:id", middlewares.AuthMiddleware(), controllers.FindUsersById) 

	router.PUT("/api/users", middlewares.AuthMiddleware(), controllers.UpdateUser)

	router.DELETE("/api/user", middlewares.AuthMiddleware(), controllers.DeleteUser)

	return router
}
