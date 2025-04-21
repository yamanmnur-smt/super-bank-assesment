package models

import (
	"gorm.io/gorm"
)

type TermDepositsTypestatus string

const (
	TERM_ACTIVE TermDepositsTypestatus = "ACTIVE"
	TERM_CLOSED TermDepositsTypestatus = "CLOSED"
)

type TermDepositsTypes struct {
	gorm.Model
	Name          string                 `gorm:"type:varchar(100);not null"`
	InterestRate  float64                `gorm:"type:decimal(5,2);not null"`
	MinAmount     float64                `gorm:"type:decimal(20,2);not null"`
	MaxAmount     float64                `gorm:"type:decimal(20,2);not null"`
	TermDays      uint                   `gorm:"not null"`
	EffectiveDate string                 `gorm:"type:date;not null"`
	Status        TermDepositsTypestatus `gorm:"column:status" sql:"type:enum('active','closed');default:'active'"`
}
