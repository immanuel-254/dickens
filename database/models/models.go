// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0

package models

import (
	"database/sql"
)

type Blog struct {
	ID        int64         `json:"id"`
	UserID    sql.NullInt64 `json:"user_id"`
	Title     string        `json:"title"`
	Body      string        `json:"body"`
	CreatedAt sql.NullTime  `json:"created_at"`
	UpdatedAt sql.NullTime  `json:"updated_at"`
}

type Category struct {
	ID        int64         `json:"id"`
	UserID    sql.NullInt64 `json:"user_id"`
	Name      string        `json:"name"`
	CreatedAt sql.NullTime  `json:"created_at"`
	UpdatedAt sql.NullTime  `json:"updated_at"`
}

type CategoryBlog struct {
	CategoryID int64        `json:"category_id"`
	BlogID     int64        `json:"blog_id"`
	CreatedAt  sql.NullTime `json:"created_at"`
	UpdatedAt  sql.NullTime `json:"updated_at"`
}

type Log struct {
	ID        int64        `json:"id"`
	DbTable   string       `json:"db_table"`
	Action    string       `json:"action"`
	ObjectID  int64        `json:"object_id"`
	CreatedAt sql.NullTime `json:"created_at"`
	UpdatedAt sql.NullTime `json:"updated_at"`
}

type Profile struct {
	ID        int64          `json:"id"`
	UserID    sql.NullInt64  `json:"user_id"`
	Username  string         `json:"username"`
	Image     sql.NullString `json:"image"`
	Bio       sql.NullString `json:"bio"`
	CreatedAt sql.NullTime   `json:"created_at"`
	UpdatedAt sql.NullTime   `json:"updated_at"`
}

type User struct {
	ID        int64        `json:"id"`
	Surname   string       `json:"surname"`
	FirstName string       `json:"first_name"`
	LastName  string       `json:"last_name"`
	Password  string       `json:"password"`
	Email     string       `json:"email"`
	CreatedAt sql.NullTime `json:"created_at"`
	UpdatedAt sql.NullTime `json:"updated_at"`
}