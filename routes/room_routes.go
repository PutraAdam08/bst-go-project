package routes

import (
	"BSTproject.com/controller"
	"BSTproject.com/middleware"

	"github.com/gin-gonic/gin"
)

func RoomRoutes(router *gin.Engine, roomController *controller.RoomController, jwtService middleware.JWTService) {
	roomRoutes := router.Group("/rooms")
	{
		roomRoutes.GET("", roomController.GetAll)
		roomRoutes.GET("/:id", roomController.GetByID)
		roomRoutes.POST("", middleware.Authenticate(jwtService), roomController.Create)
		roomRoutes.PUT("/:id", middleware.Authenticate(jwtService), roomController.Update)
		roomRoutes.DELETE("/:id", middleware.Authenticate(jwtService), roomController.Delete)
	}
}
