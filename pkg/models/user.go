package models

import "time"

type User struct {
	ID        int64
	FirstName string
	LastName  string
	Username  string
	Email     string
	Password  []byte
	Created   time.Time
	Active    bool
}
