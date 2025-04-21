package models

import "gorm.io/gorm"

type Customer struct {
	gorm.Model
	Photo          string `gorm:"type:varchar(200)"`
	Name           string `gorm:"type:varchar(200);not null"`
	Email          string `gorm:"type:varchar(200);not null"`
	PhoneNumber    string `gorm:"type:varchar(20);not null"`
	Address        string `gorm:"type:varchar(255);not null"`
	Username       string `gorm:"type:varchar(100);not null;index"`
	Gender         string `gorm:"type:varchar(10);not null"`
	AccountPurpose string `gorm:"type:varchar(100);not null"`
	SourceOfIncome string `gorm:"type:varchar(100);not null"`
	IncomePerMonth string `gorm:"type:varchar(100);not null"`
	Jobs           string `gorm:"type:varchar(100);not null"`
	Position       string `gorm:"type:varchar(100);not null"`
	Industries     string `gorm:"type:varchar(100);not null"`
	CompanyName    string `gorm:"type:varchar(100);not null"`
	AddressCompany string `gorm:"type:varchar(100);"`

	BankAccounts []BankAccount `json:"bank_accounts"`
	Pockets      []Pocket      `json:"pockets"`
}
