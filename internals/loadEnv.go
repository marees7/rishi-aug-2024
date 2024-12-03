package internals

import (
	"fmt"

	"github.com/joho/godotenv"
)

func recoverPanic() {
	if r := recover(); r != nil {
		fmt.Println("recovered from ", r)
	}
}

func LoadEnv() {
	defer recoverPanic()

	err := godotenv.Load("E:/Exercises/golang/rte_project/rishi-aug-2024/.env")
	if err != nil {
		panic(err)
	}
}
