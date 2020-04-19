package common

import (
	"fmt"
	"net/http"
	"runtime/debug"

	"github.com/supersingh05/go-authn/pkg/authn"
	"github.com/supersingh05/go-authn/pkg/models"
)

type Application struct {
	Logger Logger
	Auth   authn.Auth
	Users  models.UserDatastore
}

func (app *Application) ServerError(w http.ResponseWriter, err error) {
	trace := fmt.Sprintf("%s\n%s", err.Error(), debug.Stack())
	app.Logger.ErrorLog.Println(trace)

	http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
}

func (app *Application) ClientError(w http.ResponseWriter, status int) {
	http.Error(w, http.StatusText(status), status)
}

func (app *Application) NotFound(w http.ResponseWriter) {
	app.ClientError(w, http.StatusNotFound)
}
