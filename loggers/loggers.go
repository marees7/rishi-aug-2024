package loggers

import (
	"fmt"
	"log"
	"os"
)

var (
	WarningLog *log.Logger
	InfoLog    *log.Logger
	ErrorLog   *log.Logger
)

func OpenLog() {
	file, err := os.OpenFile(os.Getenv("file_name"), os.O_CREATE|os.O_APPEND|os.O_RDWR, 0666)
	if err != nil {
		fmt.Println(err)
		return
	}

	InfoLog = log.New(file, "INFO", log.Ldate|log.Ltime|log.Lshortfile)
	WarningLog = log.New(file, "WARNING", log.Ldate|log.Ltime|log.Lshortfile)
	ErrorLog = log.New(file, "ERROR", log.Ldate|log.Ltime|log.Lshortfile)
}
