package inputvalidators

import (
	"github.com/supersingh05/go-authn/cmd/web/requests"
	"github.com/supersingh05/go-authn/cmd/web/responses"
)

func ValidateSignup(request requests.SignupRequest) []responses.Error {
	var errors []responses.Error
	if len(request.Email) == 0 {
		errors = append(errors, responses.Error{
			Message: "field cannot be emtpy",
			Field:   "email",
		})
	}
	if len(request.Password) == 0 {
		errors = append(errors, responses.Error{
			Message: "field cannot be emtpy",
			Field:   "password",
		})
	}

	return errors
}
