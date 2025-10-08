package seeder

import (
	"fmt"

	"BSTproject.com/model"
	"BSTproject.com/utils/auth"
	db "BSTproject.com/utils/database"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()
	db := db.DBconnect()

	db.AutoMigrate(&model.User{}, &model.Room{}, &model.Booking{})

	pwd := "12345678"
	pwdHash, err := auth.HashAndSalt(pwd)
	if err != nil {
		fmt.Println("cannot hash password")
		return
	}

	var users []model.User
	users = append(users, model.User{
		Name:         "Admin1",
		PasswordHash: pwdHash,
		Email:        "admin1@gmail.com",
		IsAdmin:      true,
	})
	users = append(users, model.User{
		Name:         "Budi",
		PasswordHash: pwdHash,
		Email:        "budi@gmail.com",
		IsAdmin:      false,
	})

	db.Create(&users)

}
