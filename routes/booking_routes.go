package routes

import (
	"BSTproject.com/controller"
	"BSTproject.com/middleware"

	"github.com/gin-gonic/gin"
)

func BookingRoutes(router *gin.Engine, bookingController *controller.BookingController, jwtService middleware.JWTService) {
	bookingRoutes := router.Group("/bookings")
	{
		// Protected routes
		bookingRoutes.Use(middleware.Authenticate(jwtService))
		{
			bookingRoutes.GET("", bookingController.GetAllBookings)
			bookingRoutes.GET("/:id", bookingController.GetBookingByID)
			bookingRoutes.POST("", bookingController.CreateBooking)
			bookingRoutes.PUT("/:id", bookingController.UpdateBooking)
			bookingRoutes.DELETE("/:id", bookingController.DeleteBooking)
			bookingRoutes.PATCH("/:id/status", bookingController.UpdateBookingStatus)
		}
	}
}
