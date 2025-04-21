package db

import (
	"yamanmnur/simple-dashboard/internal/models"

	"gorm.io/gorm"
)

func Init(db *gorm.DB) *gorm.DB {

	db.AutoMigrate(&models.User{})
	db.AutoMigrate(&models.Customer{})
	db.AutoMigrate(&models.BankAccount{})
	db.AutoMigrate(&models.Pocket{})
	db.AutoMigrate(&models.TermDepositsTypes{})
	db.AutoMigrate(&models.TermDeposit{})
	db.AutoMigrate(&models.AccountTransactions{})

	return db
}
