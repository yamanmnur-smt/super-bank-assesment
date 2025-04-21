package models

import (
	"gorm.io/gorm"
)

type Pocket struct {
	gorm.Model
	CustomerID uint64   `gorm:"not null;index"`
	Name       string   `gorm:"type:varchar(100);not null"`
	Balance    float64  `gorm:"type:decimal(15,2);default:0.00"`
	Currency   string   `gorm:"type:varchar(50);not null"`
	Customer   Customer `gorm:"foreignKey:CustomerID"`
}
