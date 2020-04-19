package handler

import (
	"net/http"

	"encoding/json"

	"github.com/supersingh05/go-authn/cmd/web/requests"
	"github.com/supersingh05/go-authn/cmd/web/responses"
	"github.com/supersingh05/go-authn/pkg/common"
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
		s.ClientError(rw, http.StatusBadRequest)
		s.Logger.ErrorLog.Println(err)
		return
	}
	tokenString, err := s.Auth.CreateToken(loginreq.Username)
	if err != nil {
		s.ClientError(rw, http.StatusInternalServerError)
	}

	json.NewEncoder(rw).Encode(responses.LoginResponse{
		Token: tokenString,
	})
}
