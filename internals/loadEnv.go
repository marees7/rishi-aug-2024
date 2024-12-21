package internals

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/joho/godotenv"
)

//handle panic if occurred
func recoverPanic() {
	if r := recover(); r != nil {
		fmt.Println("recovered from ", r)
	}
}

//Load the env file
func LoadEnv() {
	defer recoverPanic()

	//get the working directory file path
	wd, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	//loads the env file
	if err = godotenv.Load(filepath.Join(filepath.Dir(wd), ".env")); err != nil {
		panic(err)
	}
}
