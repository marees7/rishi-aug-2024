package internals

import (
	"os"
	"path/filepath"

	"github.com/joho/godotenv"
)

// Load the env file
func LoadEnv() {
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
