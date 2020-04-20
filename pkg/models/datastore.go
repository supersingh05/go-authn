package models

type UserDatastore interface {
	Insert(firstname, lastname, email, password string) error

	Authenticate(email, password string) (int, error)

	Get(id int) (*User, error)
}
