package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model        // gorm default record (ID, CreatedAt, UpdatedAt, DeletedAt)
	Username   string `json:"username" gorm:"type:varchar(100);unique;not null"`
	Email      string `json:"email" gorm:"type:varchar(100);unique;not null"`
	Password   string `json:"password" gorm:"type:varchar(100); not null"`
	IsAdmin    bool   `json:"isAdmin" gorm:"default:false"`
}
