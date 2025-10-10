package main

import (
	"BSTproject.com/controller"
	"BSTproject.com/docs"
	"BSTproject.com/middleware"
	"BSTproject.com/repository"
	"BSTproject.com/routes"
	"BSTproject.com/service"
	db "BSTproject.com/utils/database"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"log"
	"os"
)

// @title           Swagger Example API
// @version         1.0
// @description     This is a sample server celler server.
// @termsOfService  http://swagger.io/terms/
// @securityDefinitions.apikey BearerToken
// @in header
// @name Authorization

// @contact.name   API Support
// @contact.url    http://www.swagger.io/support
// @contact.email  support@swagger.io

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @BasePath  /v2
func main() {
	// Load environment variables
	if err := godotenv.Load(); err != nil {
		log.Println("Warning: .env file not found")
	}

	// Database connection
	dbConn := db.DBconnect()

	// Initialize JWT service
	jwtSvc := service.NewJWTService()

	// User module setup
	userRepo := repository.NewUserRepository(dbConn)
	userSvc := service.NewUserService(jwtSvc, userRepo)
	userCtrl := controller.NewUserController(userSvc)

	// Booking module setup
	bookingRepo := repository.NewBookingRepository(dbConn)
	bookingSvc := service.NewBookingService(bookingRepo)
	bookingCtrl := controller.NewBookingController(bookingSvc)

	roomRepo := repository.NewRoomRepository(dbConn)
	roomSvc := service.NewRoomService(roomRepo)
	roomCtrl := controller.NewRoomController(roomSvc)

	// Setup Gin router
	r := gin.Default()
	r.Use(gin.Recovery())
	r.Use(middleware.CustomLogger())

	// Swagger setup
	docs.SwaggerInfo.BasePath = os.Getenv("SWAGGER_BASE_PATH")
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// Routes
	routes.UserRoutes(r, userCtrl, jwtSvc)
	routes.AdminRoutes(r, userCtrl, jwtSvc)
	routes.BookingRoutes(r, bookingCtrl, jwtSvc)
	routes.RoomRoutes(r, roomCtrl, jwtSvc)

	// Run server
	if err := r.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}
