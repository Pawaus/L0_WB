package models

import (
	"gorm.io/datatypes"
)

type Orders struct {
	Uid  string `gorm:"primarykey"`
	Data datatypes.JSON
}
