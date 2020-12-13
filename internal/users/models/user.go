package models

import "time"

type User struct {
	ID        string
	Login     string
	Email     string
	Name      string
	Password  string
	CreatedAt time.Time
	UpdatedAt time.Time
}
