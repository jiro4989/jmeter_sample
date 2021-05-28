package main

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

type Client struct {
	db *sql.DB
}

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

var client Client

func initDB() error {
	user := os.Getenv("DB_USERNAME")
	password := os.Getenv("DB_PASSWORD")
	host := os.Getenv("DB_HOSTNAME")
	port := os.Getenv("DB_PORT")
	dbname := os.Getenv("DB_DATABASE")
	dbSrc := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=true", user, password, host, port, dbname)

	db, err := sql.Open("mysql", dbSrc)
	if err != nil {
		return err
	}
	client.db = db

	return nil
}

func (c *Client) Users() ([]User, error) {
	rows, err := c.db.Query("SELECT id, name, age FROM users")
	if err != nil {
		return nil, err
	}

	var users []User
	for rows.Next() {
		var u User
		if err := rows.Scan(&u.ID, &u.Name, &u.Age); err != nil {
			return nil, err
		}
		users = append(users, u)
	}
	return users, nil
}

func (c *Client) User(userID string) (*User, error) {
	stmt, err := c.db.Prepare("SELECT id, name, age FROM users WHERE id = ?")
	if err != nil {
		return nil, err
	}

	var u User
	if err := stmt.QueryRow(userID).Scan(&u.ID, &u.Name, &u.Age); err != nil {
		return nil, err
	}
	return &u, nil
}

func (c *Client) Hobbies(userID string) ([]Hobby, error) {
	stmt, err := c.db.Prepare("SELECT id, user_id, name FROM hobbies WHERE user_id = ?")
	if err != nil {
		return nil, err
	}

	rows, err := stmt.Query(userID)
	if err != nil {
		return nil, err
	}

	var hobbies []Hobby
	for rows.Next() {
		var h Hobby
		if err := rows.Scan(&h.ID, &h.UserID, &h.Name); err != nil {
			return nil, err
		}
		hobbies = append(hobbies, h)
	}

	return hobbies, nil
}
