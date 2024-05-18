package main

import (
	importmovies "github.com/gabszero/recommendation-system-go/internal/import_movies"
	"github.com/gabszero/recommendation-system-go/internal/infra"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}
	infra.InitDatabase()

	importmovies.ImportMovie()
}
