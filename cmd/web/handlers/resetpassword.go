package handler

import (
	"net/http"

	"github.com/supersingh05/go-authn/pkg/common"
)

type ResetPasswordHandler struct {
	common.Application
}

func NewResetPasswordHandler(app common.Application) http.Handler {
	return &ResetPasswordHandler{app}
}

func (s *ResetPasswordHandler) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

}
