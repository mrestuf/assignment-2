package config

import (
	"assignment2/httpserver/repositories/models"
	"fmt"

	"github.com/jinzhu/gorm"
)

const (
	host     = "localhost"
	port     = "8080"
	user     = "postgres"
	password = "superuser"
	dbname   = "orders_by"
)

func ConnectPostgresGORM() (*gorm.DB, error) {
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	db, err := gorm.Open("postgres", dsn)
	if err != nil {
		return nil, err
	}

	db.Debug().AutoMigrate(models.Order{}, models.Item{})

	return db, nil
}
