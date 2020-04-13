package common

import "log"

type Logger struct {
	ErrorLog *log.Logger
	InfoLog  *log.Logger
}
