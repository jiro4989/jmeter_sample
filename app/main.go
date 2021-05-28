package main

import (
	"database/sql"
	"fmt"
	"net/http"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type User struct {
	ID   string
	Name string
	Age  int
}

type Hobby struct {
	ID     string
	UserID string
	Name   string
}

type ErrorResp struct {
	Message string
}

var db *sql.DB

func main() {
	user := os.Getenv("DB_USERNAME")
	password := os.Getenv("DB_PASSWORD")
	host := os.Getenv("DB_HOSTNAME")
	port := os.Getenv("DB_PORT")
	dbname := os.Getenv("DB_DATABASE")
	dbSrc := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=true", user, password, host, port, dbname)

	var err error
	db, err = sql.Open("mysql", dbSrc)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/users", apiGetUsers)
	e.GET("/users/:id", apiGetUsersUserID)
	e.GET("/users/:id/hobbies", apiGetUsersUserIDHobbies)

	e.Logger.Fatal(e.Start(":1323"))
}

func apiGetUsers(c echo.Context) error {
	rows, err := db.Query("SELECT id, name, age FROM users")
	if err != nil {
		return err
	}

	var users []User
	for rows.Next() {
		var u User
		if err := rows.Scan(&u.ID, &u.Name, &u.Age); err != nil {
			return err
		}
		users = append(users, u)
	}

	return c.JSON(http.StatusOK, users)
}

func apiGetUsersUserID(c echo.Context) error {
	userID := c.Param("id")

	stmt, err := db.Prepare("SELECT id, name, age FROM users WHERE id = ?")
	if err != nil {
		return err
	}

	var u User
	if err := stmt.QueryRow(userID).Scan(&u.ID, &u.Name, &u.Age); err != nil {
		return err
	}

	return c.JSON(http.StatusOK, u)
}

func apiGetUsersUserIDHobbies(c echo.Context) error {
	userID := c.Param("id")

	stmt, err := db.Prepare("SELECT id, user_id, name FROM hobbies WHERE user_id = ?")
	if err != nil {
		return err
	}

	rows, err := stmt.Query(userID)
	if err != nil {
		return err
	}

	var hobbies []Hobby
	for rows.Next() {
		var h Hobby
		if err := rows.Scan(&h.ID, &h.UserID, &h.Name); err != nil {
			return err
		}
		hobbies = append(hobbies, h)
	}

	return c.JSON(http.StatusOK, hobbies)
}
