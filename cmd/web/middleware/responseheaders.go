package middleware

import (
	"net/http"

	"github.com/supersingh05/go-authn/pkg/common"
)

type ResponseHeadersMiddleware struct {
	next http.Handler
	app  common.Application
}

func NewResponseHeadersMiddleware(app common.Application, next http.Handler) http.Handler {

	return SecureHeadersMiddleware{
		next: next,
		app:  app,
	}
}

func (s ResponseHeadersMiddleware) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	rw.Header().Set("Content-Type", "application/json")
	s.next.ServeHTTP(rw, r)
}
