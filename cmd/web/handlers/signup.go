package handler

import (
	"net/http"

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

}
