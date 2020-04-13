package middleware

import (
	"net/http"

	"github.com/supersingh05/go-authn/pkg/common"
)

type LogRequestMiddleware struct {
	next http.Handler
	app  common.Application
}

func NewLogRequestMiddleware(app common.Application, next http.Handler) http.Handler {

	return LogRequestMiddleware{
		next: next,
		app:  app,
	}
}

func (l LogRequestMiddleware) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	l.app.Logger.InfoLog.Printf("%s - %s %s %s", r.RemoteAddr, r.Proto, r.Method, r.URL.RequestURI())
	l.next.ServeHTTP(w, r)
}
