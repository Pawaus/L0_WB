package database

import (
	"l0_wb/internal/models"
	"l0_wb/internal/pkg/postgres"

	"gorm.io/datatypes"
)

type Database struct {
	*postgres.DB
}

func NewDatabase(host, port, user, password, dbname string) *Database {
	pg_db := postgres.NewDBConnect(host, port, user, password, dbname)
	if pg_db != nil {
		db := Database{pg_db}
		return &db
	} else {
		return nil
	}

}

func (d *Database) CreateTable() {
	d.AutoMigrate(models.Orders{})
}

func (d *Database) InsertData(uid, data string) {
	order := models.Orders{}
	order.Data = datatypes.JSON(data)
	order.Uid = uid
	d.Create(&order)
}

func (d *Database) GetOrderById(uid string) string {
	order := models.Orders{Uid: uid}
	d.First(&order)
	return string(order.Data)
}
