package middleware

import (
	"fmt"
	"net/http"

	"github.com/supersingh05/go-authn/pkg/common"
)

type RecoverPanicMiddleware struct {
	next http.Handler
	app  common.Application
}

func NewRecoverPanicMiddleware(app common.Application, next http.Handler) http.Handler {

	return RecoverPanicMiddleware{
		next: next,
		app:  app,
	}
}

func (rp RecoverPanicMiddleware) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	defer func() {
		if err := recover(); err != nil {
			rw.Header().Set("Connection", "close")
			rp.app.ServerError(rw, fmt.Errorf("%s", err))

		}
	}()
	rp.next.ServeHTTP(rw, r)
}
