package main

import (
	"net/http"

	handler "github.com/supersingh05/go-authn/cmd/web/handlers"
	"github.com/supersingh05/go-authn/pkg/common"
	"github.com/supersingh05/go-authn/pkg/config"
)

func routes(cfg config.Config, app common.Application) http.Handler {
	mux := http.NewServeMux()

	// routes
	mux.Handle("/health", handler.NewHealthHandler(app))
	mux.Handle("/login", handler.NewLoginHandler(app))
	mux.Handle("/signup", handler.NewSignupHandler(app))
	mux.Handle("/resetpassword", handler.NewResetPasswordHandler(app))

	// file server
	fileServer := http.FileServer(http.Dir(cfg.StaticDir))
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))
	return mux
	// return app.recoverPanic(app.logRequest(secureHeaders(mux)))
}
