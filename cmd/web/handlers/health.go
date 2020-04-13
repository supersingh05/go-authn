package handler

import (
	"net/http"

	"github.com/supersingh05/go-authn/pkg/common"
)

type HealthHandler struct {
	common.Application
}

func NewHealthHandler(app common.Application) http.Handler {
	return HealthHandler{app}
}

func (s HealthHandler) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

}
