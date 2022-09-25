package config

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)



// cara singleton
// var (
// 	DB *sql.DB
// )

// func ConnectPostgres() error {
// 	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

// 	db, err := sql.Open("postgres", dsn)
// 	if err != nil {
// 		return err
// 	}

// 	err = db.Ping()
// 	if err != nil {
// 		return err
// 	}

// 	DB = db

// 	return nil
// }

// func GetDB() *sql.DB {
// 	return DB
// }

func ConnectPostgres() (*sql.DB, error) {
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}
	return db, nil
}
