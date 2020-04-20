package middleware

import (
	"net/http"

	"github.com/supersingh05/go-authn/pkg/common"
)

type SetMaxBodySizeMiddleware struct {
	next http.Handler
	app  common.Application
}

func NewSetMaxBodySizeMiddleware(app common.Application, next http.Handler) http.Handler {

	return LogRequestMiddleware{
		next: next,
		app:  app,
	}
}

func (l SetMaxBodySizeMiddleware) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	r.Body = http.MaxBytesReader(rw, r.Body, 1048576)
	l.next.ServeHTTP(rw, r)
}
