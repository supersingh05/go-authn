package handler

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/supersingh05/go-authn/cmd/web/inputvalidators"
	"github.com/supersingh05/go-authn/cmd/web/requests"
	"github.com/supersingh05/go-authn/cmd/web/responses"
	"github.com/supersingh05/go-authn/pkg/common"
	"github.com/supersingh05/go-authn/pkg/models"
)

type SignupHandler struct {
	common.Application
}

func NewSignupHandler(app common.Application) http.Handler {
	return &SignupHandler{app}
}

func (s *SignupHandler) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	signupreq := requests.SignupRequest{}
	err := json.NewDecoder(r.Body).Decode(&signupreq)
	if err != nil {
		s.ClientError(rw, http.StatusInternalServerError)
		return
	}
	errorslice := inputvalidators.ValidateSignup(signupreq)

	if len(errorslice) > 0 {
		errorResponse := responses.Errors{
			Errors: errorslice,
		}
		errorBytes, err := json.Marshal(errorResponse)
		var retString string
		if err == nil {
			retString = string(errorBytes)
		}
		http.Error(rw, retString, http.StatusBadRequest)
		return
	}

	err = s.Users.Insert(signupreq.FirstName, signupreq.LastName, signupreq.Email, signupreq.Password)
	if err != nil {
		if errors.Is(err, models.ErrDuplicateEmail) {
			errorResponse := responses.Errors{
				Errors: []responses.Error{{
					Message: "already used",
					Field:   "email",
				}},
			}
			errorBytes, err := json.Marshal(errorResponse)
			var retString string
			if err == nil {
				retString = string(errorBytes)
			}
			http.Error(rw, retString, http.StatusBadRequest)
		} else {
			s.ServerError(rw, err)
		}
		return
	}
}
