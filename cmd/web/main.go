package main

import (
	"log"
	"os"

	"github.com/supersingh05/go-authn/pkg/common"
	"github.com/supersingh05/go-authn/pkg/config"
)

type Application struct {
	logger common.Logger
}

func main() {
	cfg := config.ParseConfig()
	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	app := &Application{
		logger: common.Logger {
			InfoLog: infoLog,
			ErrorLog: errorLog,
		}
	}

	srv := &http.Server{
		Addr:     cfg.Addr,
		ErrorLog: errorLog,
		Handler:  app.routes(cfg),
	}

	infoLog.Printf("Starting server on %s", cfg.Addr)
	// errorLog.Fatal(srv.ListenAndServe())

}
