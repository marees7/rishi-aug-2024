package loggers

import (
	"log"
	"os"
	"path/filepath"
)

var (
	Warn  *log.Logger
	Info  *log.Logger
	Error *log.Logger
)

// creates a new log file or opens if file already exists
func OpenLog() {
	//get the working directory file path
	wd, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	//creates a new log file or opens if file already exists
	file, err := os.OpenFile(filepath.Join(filepath.Dir(wd), os.Getenv("FILE_NAME")), os.O_CREATE|os.O_APPEND|os.O_RDWR, 0666)
	if err != nil {
		panic(err)
	}

	//set the log prefix and flags
	Info = log.New(file, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
	Warn = log.New(file, "WARNING: ", log.Ldate|log.Ltime|log.Lshortfile)
	Error = log.New(file, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)
}
