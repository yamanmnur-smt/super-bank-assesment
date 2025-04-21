package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username string `gorm:"type:varchar(100);not null;index"`
	Password string `gorm:"type:varchar(200);not null"`
	Name     string `gorm:"type:varchar(200);not null"`
}
