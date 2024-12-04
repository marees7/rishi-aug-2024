package internals

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/joho/godotenv"
)

func recoverPanic() {
	if r := recover(); r != nil {
		fmt.Println("recovered from ", r)
	}
}

func LoadEnv() {
	defer recoverPanic()

	wd, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	if err = godotenv.Load(filepath.Join(filepath.Dir(wd), ".env")); err != nil {
		panic(err)
	}
}
