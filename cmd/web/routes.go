package main

import (
	"net/http"

	handler "github.com/supersingh05/go-authn/cmd/web/handlers"
	"github.com/supersingh05/go-authn/cmd/web/middleware"
	"github.com/supersingh05/go-authn/pkg/common"
	"github.com/supersingh05/go-authn/pkg/config"
)

func routes(cfg config.Config, app common.Application) http.Handler {
	mux := http.NewServeMux()

	// routes
	mux.Handle("/health", handler.NewHealthHandler(app))

	onlyAllowPostLogin := middleware.NewMethodsAllowedMiddleware(app, []string{http.MethodPost}, handler.NewLoginHandler(app))
	mux.Handle("/login", onlyAllowPostLogin)

	mux.Handle("/signup", handler.NewSignupHandler(app))

	onlyAllowPostReset := middleware.NewMethodsAllowedMiddleware(app, []string{http.MethodPost}, handler.NewResetPasswordHandler(app))
	mux.Handle("/resetpassword", onlyAllowPostReset)

	// file server
	fileServer := http.FileServer(http.Dir(cfg.StaticDir))
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	//universal middleware
	secureHeaders := middleware.NewSecureHeadersMiddleware(app, mux)
	logMiddleware := middleware.NewLogRequestMiddleware(app, secureHeaders)
	recoverPanic := middleware.NewSecureHeadersMiddleware(app, logMiddleware)

	return recoverPanic
}
