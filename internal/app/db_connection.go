package app

import (
	"fmt"

	"github.com/guzmanmo/go-imdb/internal/app/repository"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DbConnection *gorm.DB

func SetUpConnection() *gorm.DB {
	connectToDatabase()
	createTables()
	return DbConnection
}

func connectToDatabase() {
	dsn := "host=localhost user=postgres password=test1234 dbname=imdb port=5432 sslmode=disable"
	var err error
	DbConnection, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
}

func createTables() {
	err := DbConnection.AutoMigrate(repository.Movie{}, repository.Actor{}, repository.Evaluation{})
	if err != nil {
		fmt.Println(err)
	}

}
