package mock

import (
	"github.com/supersingh05/go-authn/pkg/models"
)

type UserModel struct {
}

// We'll use the Insert method to add a new record to the users table.
func (m *UserModel) Insert(firstname, lastname, email, password string) error {
	return nil
}

// We'll use the Authenticate method to verify whether a user exists with // the provided email address and password. This will return the relevant // user ID if they do.
func (m *UserModel) Authenticate(email, password string) (int, error) {
	return 0, nil
}

// We'll use the Get method to fetch details for a specific user based // on their user ID.
func (m *UserModel) Get(id int) (*models.User, error) {
	return nil, nil
}
