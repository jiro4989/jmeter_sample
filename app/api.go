package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func apiGetUsers(c echo.Context) error {
	users, err := client.Users()
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, users)
}

func apiGetUsersUserID(c echo.Context) error {
	userID := c.Param("id")

	u, err := client.User(userID)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, u)
}

func apiGetUsersUserIDHobbies(c echo.Context) error {
	userID := c.Param("id")

	hobbies, err := client.Hobbies(userID)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, hobbies)
}
