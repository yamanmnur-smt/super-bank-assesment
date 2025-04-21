package models

import "gorm.io/gorm"

type BankAccountStatus string

const (
	ACTIVE_BANK_ACCOUNT  BankAccountStatus = "ACTIVE"
	BLOCKED_BANK_ACCOUNT BankAccountStatus = "BLOCKED"
	CLOSED_BANK_ACCOUNT  BankAccountStatus = "CLOSED"
)

type BankAccount struct {
	gorm.Model
	CustomerID     uint              `gorm:"not null;index"`
	CardNumber     string            `gorm:"type:varchar(50);uniqueIndex;not null"`
	AccountNumber  string            `gorm:"type:varchar(50);uniqueIndex;not null"`
	Balance        float64           `gorm:"type:decimal(20,2);default:0.00"`
	AccountType    string            `gorm:"type:varchar(50);not null"`
	Cvc            string            `gorm:"type:varchar(10);not null"`
	ExpirationDate string            `gorm:"type:date;not null"`
	Status         BankAccountStatus `gorm:"column:status" sql:"type:enum('ACTIVE_BANK_ACCOUNT','BLOCKED_BANK_ACCOUNT','CLOSED_BANK_ACCOUNT');default:'ACTIVE_BANK_ACCOUNT'"`
	Customer       Customer          `gorm:"foreignKey:CustomerID"`
	TermDeposit    []TermDeposit     `json:"term_deposit"`
}
