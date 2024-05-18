package main

import (
	"github.com/gabszero/recommendation-system-go/internal/imports"
	"github.com/gabszero/recommendation-system-go/internal/infra"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}

	infra.InitDatabase()
	imports.ImportMovie()
	imports.ImportUsers()
}
