package user

import "time"

type User struct {
	ID        string
	Name      string
	Email     string
	Password  string
	createdAt time.Time
	updatedAt time.Time
}
