package middleware

import (
	"net/http"

	"github.com/golang/gddo/httputil/header"
	"github.com/supersingh05/go-authn/pkg/common"
)

type EnforceJsonMiddleware struct {
	next http.Handler
	app  common.Application
}

func NewEnforceJsonMiddleware(app common.Application, next http.Handler) http.Handler {

	return EnforceJsonMiddleware{
		next: next,
		app:  app,
	}
}

func (l EnforceJsonMiddleware) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	if r.Header.Get("Content-Type") != "" {
		value, _ := header.ParseValueAndParams(r.Header, "Content-Type")
		if value != "application/json" {
			l.app.ClientError(rw, http.StatusUnsupportedMediaType)
			return
		}
	}
	l.next.ServeHTTP(rw, r)
}
