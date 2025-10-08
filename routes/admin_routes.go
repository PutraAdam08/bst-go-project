package routes

import (
	"BSTproject.com/controller"
	"BSTproject.com/middleware"

	"github.com/gin-gonic/gin"
)

func AdminRoutes(router *gin.Engine, userController *controller.UserController, jwtService middleware.JWTService) {
	userRoutes := router.Group("/admin")
	{
		userRoutes.POST("/login", userController.AdminLogin)
		userRoutes.GET("", middleware.Authenticate(jwtService), userController.GetUser)
	}
}
