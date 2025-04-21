package models

import "gorm.io/gorm"

type TransactionStatus string

const (
	COMPLETED TransactionStatus = "COMPLETED"
	PENDING   TransactionStatus = "PENDING"
	FAILED    TransactionStatus = "FAILED"
)

type TransactionType string

const (
	DEPOSITS TransactionType = "DEPOSITS"
	POCKETS  TransactionType = "POCKETS"
	TRANSFER TransactionType = "TRANSFER"
	DEBIT    TransactionType = "DEBIT"
)

type AccountTransactions struct {
	gorm.Model
	CustomerID        uint `gorm:"not null"`
	Customer          Customer
	AccountNumber     string            `gorm:"type:varchar(50);not null"`
	TransactionType   TransactionType   `gorm:"column:transaction_type" sql:"type:enum('DEPOSITS','POCKETS','TRANSFER','DEBIT');default:'DEPOSITS'"`
	TransactionDate   string            `gorm:"type:date;not null"`
	TransactionAmount float64           `gorm:"type:decimal(20,2);not null"`
	TransactionStatus string            `gorm:"type:varchar(50);not null"`
	Description       string            `gorm:"type:varchar(255);not null"`
	ReferenceNumber   uint              `gorm:"not null"`
	Status            TransactionStatus `gorm:"column:status" sql:"type:enum('COMPLETED','PENDING','FAILED');default:'PENDING'"`
}
