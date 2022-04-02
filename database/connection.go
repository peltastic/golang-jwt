package database

import (
	_ "github.com/lib/pq"
	"github.com/peltastic/golang-jwt/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	dsn := "host=localhost user=postgres password=3123pex3123 dbname=postgres port=5432 sslmode=disable"
	conn, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("could not connect to database")
	}
	DB = conn

	conn.AutoMigrate(&models.User{})
}
