package routes

import (
	"github.com/gin-gonic/gin"
	"x-app/api/controllers"
)

func RegisterRoutes(router *gin.Engine) {
	api := router.Group("/api")
	{
		api.POST("/users", controllers.CreateUser)
		api.GET("/users", controllers.GetUsers)
		api.GET("/users/:id", controllers.GetUserByID)
		api.DELETE("/users/:id", controllers.DeleteUser)
	}
}
