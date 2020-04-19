package inputvalidators

import (
	"github.com/supersingh05/go-authn/cmd/web/requests"
	"github.com/supersingh05/go-authn/cmd/web/responses"
)

func ValidateSignup(request requests.SignupRequest) []responses.Error {
	errors := make([]responses.Error, 10)
	i := 0
	if len(request.Email) == 0 {
		errors[i] = responses.Error{
			Message: "field cannot be emtpy",
			Field:   "email",
		}
	}
	if len(request.Username) == 0 {
		errors[i] = responses.Error{
			Message: "field cannot be emtpy",
			Field:   "username",
		}
	}
	if len(request.Password) == 0 {
		errors[i] = responses.Error{
			Message: "field cannot be emtpy",
			Field:   "password",
		}
	}

	return errors[0 : i+1]
}
