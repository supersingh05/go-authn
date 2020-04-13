package handler

import (
	"net/http"

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

}
