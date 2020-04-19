package main

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/supersingh05/go-authn/pkg/authn"
	"github.com/supersingh05/go-authn/pkg/common"
	"github.com/supersingh05/go-authn/pkg/config"
	"github.com/supersingh05/go-authn/pkg/models/mysql"
)

func main() {
	cfg := config.ParseConfig()
	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	app := common.Application{
		Logger: common.Logger{
			InfoLog:  infoLog,
			ErrorLog: errorLog,
		},
		Auth:  authn.NewSimpleAuth([]byte("salt"), (5 * time.Minute)),
		Users: &mysql.UserModel{},
	}

	srv := &http.Server{
		Addr:     cfg.Addr,
		ErrorLog: app.Logger.ErrorLog,
		Handler:  routes(cfg, app),
	}

	infoLog.Printf("Starting server on %s", cfg.Addr)
	errorLog.Fatal(srv.ListenAndServe())

}
