package main

import (
	"first-project/handler"
	"first-project/user"
	"log"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	dsn := "host=localhost user=postgres password=qwerty dbname=crowd_funding port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err.Error())
	}

	userRepository := user.NewRepository(db)
	userService := user.NewService(userRepository)
	// userInput := user.RegisterUserInput{}
	userHandler := handler.NewUserHandler(userService)

	route := gin.Default()

	api := route.Group("/api/v1")
	api.POST("/user", userHandler.RegisterUser)

	route.Run()

}
