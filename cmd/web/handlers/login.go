package handler

import (
	"errors"
	"net/http"

	"encoding/json"

	"github.com/supersingh05/go-authn/cmd/web/requests"
	"github.com/supersingh05/go-authn/cmd/web/responses"
	"github.com/supersingh05/go-authn/pkg/common"
	"github.com/supersingh05/go-authn/pkg/models"
)

type LoginHandler struct {
	common.Application
}

func NewLoginHandler(app common.Application) http.Handler {
	return LoginHandler{app}
}

func (s LoginHandler) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	loginreq := requests.LoginRequest{}
	err := json.NewDecoder(r.Body).Decode(&loginreq)
	if err != nil {
		s.ClientError(rw, http.StatusInternalServerError)
		s.Logger.ErrorLog.Println(err)
		return
	}

	id, err := s.Users.Authenticate(loginreq.Email, loginreq.Password)
	if err != nil {
		if errors.Is(err, models.ErrInvalidCredentials) {
			s.ClientError(rw, http.StatusUnauthorized)
		} else {
			s.ServerError(rw, err)
		}
		return
	}
	tokenString, err := s.Auth.CreateToken(loginreq.Email, id)
	if err != nil {
		s.ClientError(rw, http.StatusInternalServerError)
	}

	json.NewEncoder(rw).Encode(responses.LoginResponse{
		Token: tokenString,
	})
}
