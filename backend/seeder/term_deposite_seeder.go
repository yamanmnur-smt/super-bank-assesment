package seeder

import (
	"yamanmnur/simple-dashboard/internal/models"
	"yamanmnur/simple-dashboard/pkg/db"
)

func TermDepositSeeder() {
	gormdb := db.GetDbHandler().DB
	var defaultTermDepositTypes models.TermDepositsTypes
	gormdb.Model(&models.TermDepositsTypes{}).Where("id = ?", 1).First(&defaultTermDepositTypes)
	termdeposits := []models.TermDeposit{
		{
			CustomerID:            1,
			BankAccountID:         1,
			Amount:                10000.00,
			InterestRate:          3.50,
			StartDate:             "2023-01-01",
			MaturityDate:          "2024-01-01",
			Status:                models.ACTIVE,
			ExtensionInstructions: models.PRINCIPAL_AND_INTEREST_ROLLOVER,
			TermDepositsTypeID:    1,
		},
		{
			CustomerID:            2,
			BankAccountID:         2,
			Amount:                15000.00,
			InterestRate:          4.00,
			StartDate:             "2023-02-01",
			MaturityDate:          "2024-02-01",
			Status:                models.ACTIVE,
			ExtensionInstructions: models.PRINCIPAL_AND_INTEREST_ROLLOVER,
			TermDepositsTypeID:    2,
		},
		{
			CustomerID:            3,
			BankAccountID:         3,
			Amount:                20000.00,
			InterestRate:          3.75,
			StartDate:             "2023-03-01",
			MaturityDate:          "2024-03-01",
			Status:                models.ACTIVE,
			ExtensionInstructions: models.PRINCIPAL_AND_INTEREST_ROLLOVER,
			TermDepositsTypeID:    3,
		},
		{
			CustomerID:            4,
			BankAccountID:         4,
			Amount:                25000.00,
			InterestRate:          4.25,
			StartDate:             "2023-04-01",
			MaturityDate:          "2024-04-01",
			Status:                models.ACTIVE,
			ExtensionInstructions: models.PRINCIPAL_AND_INTEREST_ROLLOVER,
			TermDepositsTypeID:    4,
		},
		{
			CustomerID:            5,
			BankAccountID:         5,
			Amount:                30000.00,
			InterestRate:          3.90,
			StartDate:             "2023-05-01",
			MaturityDate:          "2024-05-01",
			Status:                models.ACTIVE,
			ExtensionInstructions: models.PRINCIPAL_AND_INTEREST_ROLLOVER,
			TermDepositsTypeID:    5,
		},
		{
			CustomerID:            6,
			BankAccountID:         6,
			Amount:                35000.00,
			InterestRate:          4.10,
			StartDate:             "2023-06-01",
			MaturityDate:          "2024-06-01",
			Status:                models.ACTIVE,
			ExtensionInstructions: models.PRINCIPAL_AND_INTEREST_ROLLOVER,
			TermDepositsTypeID:    6,
		},
		{
			CustomerID:            7,
			BankAccountID:         7,
			Amount:                40000.00,
			InterestRate:          3.80,
			StartDate:             "2023-07-01",
			MaturityDate:          "2024-07-01",
			Status:                models.ACTIVE,
			ExtensionInstructions: models.PRINCIPAL_AND_INTEREST_ROLLOVER,
			TermDepositsTypeID:    7,
		},
		{
			CustomerID:            8,
			BankAccountID:         8,
			Amount:                45000.00,
			InterestRate:          4.20,
			StartDate:             "2023-08-01",
			MaturityDate:          "2024-08-01",
			Status:                models.ACTIVE,
			ExtensionInstructions: models.PRINCIPAL_AND_INTEREST_ROLLOVER,
			TermDepositsTypeID:    8,
		},
		{
			CustomerID:            9,
			BankAccountID:         9,
			Amount:                50000.00,
			InterestRate:          3.95,
			StartDate:             "2023-09-01",
			MaturityDate:          "2024-09-01",
			Status:                models.ACTIVE,
			ExtensionInstructions: models.PRINCIPAL_AND_INTEREST_ROLLOVER,
			TermDepositsTypeID:    9,
		},
		{
			CustomerID:            10,
			BankAccountID:         10,
			Amount:                55000.00,
			InterestRate:          4.30,
			StartDate:             "2023-10-01",
			MaturityDate:          "2024-10-01",
			Status:                models.ACTIVE,
			ExtensionInstructions: models.PRINCIPAL_AND_INTEREST_ROLLOVER,
			TermDepositsType:      defaultTermDepositTypes,
		},
	}

	var count int64
	gormdb.Model(&models.TermDeposit{}).Count(&count)
	if count == 0 {
		for _, item := range termdeposits {
			gormdb.Create(&item)
		}
	}
}
