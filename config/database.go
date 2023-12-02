package config

import (
	"os"

	"github.com/llucasramao/APIGO/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DB  *gorm.DB
	err error
)

func ConnectDatabase() {
	connection := "host=localhost user=root password=root dbname=root port=5432 sslmode=disable TimeZone=America/Sao_Paulo"
	DB, err = gorm.Open(postgres.Open(connection))
	if err != nil {
		logger.Errf("error to connect to database: %v", err)
		os.Exit(1)
	}
	DB.AutoMigrate(&models.Aluno{}) // Send data structure to database
}
