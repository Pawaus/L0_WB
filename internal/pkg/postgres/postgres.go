package postgres

import (
	"fmt"

	// "l0_wb/internal/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type DB struct {
	*gorm.DB
}

func NewDBConnect(host, port, user, password, dbname string) *DB {
	db_connect := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", host, user, password, dbname, port)

	db, err := gorm.Open(postgres.Open(db_connect), &gorm.Config{})
	if err != nil {
		fmt.Println("Error creating database connection " + err.Error())
		return nil
	}
	return &DB{db}
}
