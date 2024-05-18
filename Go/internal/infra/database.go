package infra

import (
	"fmt"
	"os"

	"github.com/gabszero/recommendation-system-go/internal/infra/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

func InitDatabase() {
	dbHost := os.Getenv("RECOMMENDATION_SYSTEM_HOST")
	dbPort := os.Getenv("RECOMMENDATION_SYSTEM_DB_PORT")
	dbUser := os.Getenv("RECOMMENDATION_SYSTEM_USER")
	dbPass := os.Getenv("RECOMMENDATION_SYSTEM_PASSWORD")
	dbName := os.Getenv("RECOMMENDATION_SYSTEM_DB")
	connectionString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", dbUser, dbPass, dbHost, dbPort, dbName)
	database, err := gorm.Open(mysql.Open(connectionString), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	db = database

	db.AutoMigrate(models.Movie{})
}
