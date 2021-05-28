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

func apiPostUsers(c echo.Context) error {
	var u User
	if err := c.Bind(&u); err != nil {
		return err
	}
	if u.Name == "" {
		return nil
	}
	if u.Age < 1 {
		return nil
	}

	user, err := client.NewUser(u.Name, u.Age)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, user)
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
