package main

import (
	"fmt"
	"log"

	"github.com/labstack/echo/v4"
)

func main() {
	db := mysql.InitDB()
	e := echo.New()

	controller := user.UserController{DB: db}

	e.GET("/user", controller.GetAllData())
	e.POST("/user", controller.CreateUser())
	e.GET("/user/:id", controller.GetSpecificUser())
	e.PUT("/user/:id", controller.UpdateUser())
	e.DELETE("/user/:id", controller.DeleteUser())

	fmt.Println("Running program ...")
	err := e.Start(":8000")
	if err != nil {
		log.Fatal(err.Error())
	}
}
