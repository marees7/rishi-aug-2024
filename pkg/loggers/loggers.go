package loggers

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
)

var (
	Warn  *log.Logger
	Info  *log.Logger
	Error *log.Logger
)

func recoverPanic() {
	if r := recover(); r != nil {
		fmt.Println("recovered from ", r)
	}
}

func OpenLog() {
	defer recoverPanic()

	wd, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	file, err := os.OpenFile(filepath.Join(filepath.Dir(wd), os.Getenv("FILE_NAME")), os.O_CREATE|os.O_APPEND|os.O_RDWR, 0666)
	if err != nil {
		panic(err)
	}

	Info = log.New(file, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
	Warn = log.New(file, "WARNING: ", log.Ldate|log.Ltime|log.Lshortfile)
	Error = log.New(file, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)
}
