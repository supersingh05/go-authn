package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/supersingh05/go-authn/pkg/authn"
	"github.com/supersingh05/go-authn/pkg/common"
	"github.com/supersingh05/go-authn/pkg/config"
	"github.com/supersingh05/go-authn/pkg/models/mysql"
)

func main() {
	cfg := config.ParseConfig()
	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)
	db, err := openDB(cfg.Dsn)
	if err != nil {
		errorLog.Fatal(err)
	}
	defer db.Close()
	app := common.Application{
		Logger: common.Logger{
			InfoLog:  infoLog,
			ErrorLog: errorLog,
		},
		Auth:  authn.NewSimpleAuth([]byte("salt"), (5 * time.Minute)),
		Users: &mysql.UserModel{DB: db},
	}

	srv := &http.Server{
		Addr:     cfg.Addr,
		ErrorLog: app.Logger.ErrorLog,
		Handler:  routes(cfg, app),
	}

	infoLog.Printf("Starting server on %s", cfg.Addr)
	errorLog.Fatal(srv.ListenAndServe())

}

func openDB(dsn string) (*sql.DB, error) {
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}
	if err = db.Ping(); err != nil {
		return nil, err
	}
	return db, nil
}
