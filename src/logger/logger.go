package logger

import (
	"io"
	"log"
)

var (
	//Info info logger
	Info *log.Logger
	//Debug debug logger
	Debug *log.Logger
)
//Init initializes the loggers in logger package
func Init(info, debug io.Writer) {
	Debug = log.New(debug, "DEBUG: ", log.Ldate|log.Ltime|log.Lshortfile)
	Info = log.New(info,"INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
}
