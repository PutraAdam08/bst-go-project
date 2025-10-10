package routes

import (
	"BSTproject.com/controller"
	"BSTproject.com/middleware"

	"github.com/gin-gonic/gin"
)

func UserRoutes(router *gin.Engine, userController *controller.UserController, jwtService middleware.JWTService) {
	userRoutes := router.Group("/users")
	{
		userRoutes.POST("/register", userController.Register)
		userRoutes.POST("/login", userController.Login)
		userRoutes.GET("", middleware.Authenticate(jwtService), userController.GetUser)
		userRoutes.PUT("", middleware.Authenticate(jwtService), userController.UpdateUser)
	}
}
