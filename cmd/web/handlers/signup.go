package handler

import (
	"encoding/json"
	"net/http"

	"github.com/supersingh05/go-authn/cmd/web/inputvalidators"
	"github.com/supersingh05/go-authn/cmd/web/requests"
	"github.com/supersingh05/go-authn/cmd/web/responses"
	"github.com/supersingh05/go-authn/pkg/common"
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
		s.ClientError(rw, http.StatusBadRequest)
		return
	}
	errors := inputvalidators.ValidateSignup(signupreq)
	if len(errors) > 0 {
		errorResponse := responses.Errors{
			Errors: errors,
		}
		errorBytes, err := json.Marshal(errorResponse)
		var retString string
		if err == nil {
			retString = string(errorBytes)
		}
		http.Error(rw, retString, http.StatusBadRequest)
		return
	}
	s.Users.Insert(signupreq.FirstName, signupreq.Email, signupreq.Password)

}
