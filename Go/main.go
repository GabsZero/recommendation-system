package main

import (
	"github.com/gabszero/recommendation-system-go/internals/infra"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}
	infra.InitDatabase()
}
