package main

import (
	"fmt"
	"log"

	"github.com/jackthepanda96/Belajar-Rest.git/controller/user"
	"github.com/jackthepanda96/Belajar-Rest.git/database/mysql"
	"github.com/jackthepanda96/Belajar-Rest.git/model"
	"github.com/labstack/echo/v4"
)

func main() {
	db := mysql.InitDB()
	mysql.MigrateData(db)
	e := echo.New()

	userModel := model.UserModel{DB: db}
	userController := user.UserController{Model: userModel}

	user := e.Group("/user")
	user.GET("", userController.GetAllData())
	user.POST("", userController.CreateUser())
	user.GET("/:id", userController.GetSpecificUser())
	user.PUT("/:id", userController.UpdateUser())
	user.DELETE("/:id", userController.DeleteUser())

	fmt.Println("Running program ...")
	err := e.Start(":8000")
	if err != nil {
		log.Fatal(err.Error())
	}
}
