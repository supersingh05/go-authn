package middleware

import (
	"net/http"

	"github.com/supersingh05/go-authn/pkg/common"
)

type MethodsAllowedMiddleware struct {
	next    http.Handler
	app     common.Application
	methods []string
}

func NewMethodsAllowedMiddleware(app common.Application, methods []string, next http.Handler) http.Handler {
	// MethodGet     = "GET"
	// MethodHead    = "HEAD"
	// MethodPost    = "POST"
	// MethodPut     = "PUT"
	// MethodPatch   = "PATCH" // RFC 5789
	// MethodDelete  = "DELETE"
	// MethodConnect = "CONNECT"
	// MethodOptions = "OPTIONS"
	// MethodTrace   = "TRACE"
	return MethodsAllowedMiddleware{
		next:    next,
		app:     app,
		methods: methods,
	}
}

func (l MethodsAllowedMiddleware) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	foundMethod := false
	for _, v := range l.methods {
		if r.Method == v {
			foundMethod = true
		}
	}
	if foundMethod {
		l.next.ServeHTTP(rw, r)
	} else {
		l.app.Logger.InfoLog.Printf("Method: %s, not allowed", r.Method)
		rw.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

}
