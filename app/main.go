package main

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type User struct {
	ID   string
	Name string
	Age  int
}

type ErrorResp struct {
	Message string
}

var dataUsers = map[string]User{
	"ba1ea8a5": {
		ID:   "ba1ea8a5",
		Name: "田中太郎",
		Age:  22,
	},
	"23aea8a5": {
		ID:   "23aea8a5",
		Name: "山本一郎",
		Age:  16,
	},
	"ba9f8acc": {
		ID:   "ba9f8acc",
		Name: "山田花子",
		Age:  20,
	},
	"b22adfa5": {
		ID:   "b22adfa5",
		Name: "川田圭佑",
		Age:  25,
	},
}

func main() {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/users", apiGetUsers)
	e.GET("/users/:id", apiGetUsersUserID)

	e.Logger.Fatal(e.Start(":1323"))
}

func apiGetUsers(c echo.Context) error {
	var u []User
	for _, user := range dataUsers {
		u = append(u, user)
	}
	return c.JSON(http.StatusOK, u)
}

func apiGetUsersUserID(c echo.Context) error {
	id := c.Param("id")
	if user, ok := dataUsers[id]; ok {
		return c.JSON(http.StatusOK, user)
	}
	e := ErrorResp{Message: fmt.Sprintf("%s was not found", id)}
	return c.JSON(http.StatusNotFound, e)
}
