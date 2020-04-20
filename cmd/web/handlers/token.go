package handler

import (
	"net/http"
	"strings"

	"github.com/supersingh05/go-authn/pkg/common"
)

type TokenHandler struct {
	common.Application
}

func NewTokenHandler(app common.Application) http.Handler {
	return TokenHandler{app}
}

func (s TokenHandler) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	reqToken := r.Header.Get("Authorization")
	splitToken := strings.Split(reqToken, "Bearer ")
	if len(splitToken) != 2 {
		s.ClientError(rw, http.StatusBadRequest)
		return
	}
	reqToken = strings.TrimSpace(splitToken[1])
	err := s.Auth.IsTokenValid(reqToken)
	if err != nil {
		s.ClientError(rw, http.StatusUnauthorized)
		return
	}
}
