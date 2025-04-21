package models

import (
	"gorm.io/gorm"
)

type TermStatus string

const (
	ACTIVE TermStatus = "ACTIVE"
	MATURE TermStatus = "MATURE"
	CLOSED TermStatus = "CLOSED"
)

type TermExtensionInstruction string

const (
	PRINCIPAL_ONLY_ROLLOVER         TermExtensionInstruction = "Principal only rollover"
	PRINCIPAL_AND_INTEREST_ROLLOVER TermExtensionInstruction = "Principal and interest rollover"
	NO_ROLLOVER                     TermExtensionInstruction = "No rollover"
)

type TermDeposit struct {
	gorm.Model
	CustomerID         uint              `gorm:"not null"`
	BankAccountID      uint              `gorm:"not null"`
	Amount             float64           `gorm:"type:decimal(18,2);not null"`
	InterestRate       float64           `gorm:"type:decimal(5,2);not null"`
	TermDepositsTypeID uint              `gorm:"not null"`
	TermDepositsType   TermDepositsTypes `gorm:"foreignKey:TermDepositsTypeID"`
	StartDate          string            `gorm:"type:date;not null"`
	MaturityDate       string            `gorm:"type:date;not null"`
	Customer           Customer          `gorm:"foreignKey:CustomerID"`
	BankAccount        BankAccount       `gorm:"foreignKey:BankAccountID"`

	ExtensionInstructions TermExtensionInstruction `gorm:"column:status" sql:"type:enum('PRINCIPAL_ONLY_ROLLOVER','PRINCIPAL_AND_INTEREST_ROLLOVER','NO_ROLLOVER');default:'NO_ROLLOVER'"`
	Status                TermStatus               `gorm:"column:status" sql:"type:enum('active','matured','closed');default:'active'"`
}
