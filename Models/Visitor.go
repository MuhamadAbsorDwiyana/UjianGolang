package models

import (
	"gorm.io/gorm"
)

type Visitor struct {
	gorm.Model        // gorm default record (ID, CreatedAt, UpdatedAt, DeletedAt)
	Name       string `json:"name" gorm:"type:varchar(100);not null"`
	Identity   string `json:"identity" gorm:"type:varchar(20);unique;not null"`
	Address    string `json:"address" gorm:"type:text;not null"`
	Age        uint8  `json:"age" gorm:"type:int(3);not null"`
}
