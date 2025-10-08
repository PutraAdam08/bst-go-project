package main

import (
	"log"

	"BSTproject.com/controller"
	"BSTproject.com/middleware"
	"BSTproject.com/repository"
	"BSTproject.com/routes"
	"BSTproject.com/service"
	db "BSTproject.com/utils/database"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()
	db := db.DBconnect()

	userRepo := repository.NewUserRepository(db)

	jwtSvc := service.NewJWTService()
	userSvc := service.NewUserService(jwtSvc, userRepo)

	userCtrl := controller.NewUserController(userSvc)

	r := gin.Default()
	r.Use(gin.Recovery())
	r.Use(middleware.CustomLogger())

	routes.UserRoutes(r, userCtrl, jwtSvc)

	if err := r.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}
