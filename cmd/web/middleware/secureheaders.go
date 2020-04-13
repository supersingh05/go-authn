package middleware

import (
	"net/http"

	"github.com/supersingh05/go-authn/pkg/common"
)

type SecureHeadersMiddleware struct {
	next http.Handler
	app  common.Application
}

func NewSecureHeadersMiddleware(app common.Application, next http.Handler) http.Handler {

	return SecureHeadersMiddleware{
		next: next,
		app:  app,
	}
}

func (s SecureHeadersMiddleware) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	rw.Header().Set("X-XSS-Protection", "1; mode=block")
	rw.Header().Set("X-Frame-Options", "deny")
	s.next.ServeHTTP(rw, r)
}
