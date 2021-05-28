package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	if err := initDB(); err != nil {
		panic(err)
	}

	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/users", apiGetUsers)
	e.POST("/users", apiPostUsers)
	e.GET("/users/:id", apiGetUsersUserID)
	e.GET("/users/:id/hobbies", apiGetUsersUserIDHobbies)
	e.POST("/users/:id/hobbies", apiPostHobbies)

	e.Logger.Fatal(e.Start(":1323"))
}
