package seeder

import (
	"yamanmnur/simple-dashboard/internal/models"
	"yamanmnur/simple-dashboard/pkg/db"
)

func BankAccountSeeder() {
	bankaccounts := []models.BankAccount{
		{
			CustomerID:     1,
			AccountNumber:  "1234567890",
			Balance:        1000.50,
			AccountType:    "savings",
			Cvc:            "456",
			ExpirationDate: "2026-01-15",
			CardNumber:     "9876-5432-1098-7654",
			Status:         models.ACTIVE_BANK_ACCOUNT,
		},
		{
			CustomerID:     2,
			AccountNumber:  "2345678901",
			Balance:        2500.75,
			AccountType:    "checking",
			Cvc:            "789",
			ExpirationDate: "2027-03-20",
			CardNumber:     "8765-4321-0987-6543",
			Status:         models.ACTIVE_BANK_ACCOUNT,
		},
		{
			CustomerID:     3,
			AccountNumber:  "3456789012",
			Balance:        500.00,
			AccountType:    "savings",
			Cvc:            "321",
			ExpirationDate: "2025-08-10",
			CardNumber:     "7654-3210-9876-5432",
			Status:         models.ACTIVE_BANK_ACCOUNT,
		},
		{
			CustomerID:     4,
			AccountNumber:  "4567890123",
			Balance:        3000.00,
			AccountType:    "checking",
			Cvc:            "654",
			ExpirationDate: "2026-11-05",
			CardNumber:     "6543-2109-8765-4321",
			Status:         models.ACTIVE_BANK_ACCOUNT,
		},
		{
			CustomerID:     5,
			AccountNumber:  "5678901234",
			Balance:        1500.25,
			AccountType:    "savings",
			Cvc:            "987",
			ExpirationDate: "2027-07-25",
			CardNumber:     "5432-1098-7654-3210",
			Status:         models.ACTIVE_BANK_ACCOUNT,
		},
		{
			CustomerID:     6,
			AccountNumber:  "6789012345",
			Balance:        2000.00,
			AccountType:    "checking",
			Cvc:            "123",
			ExpirationDate: "2025-09-15",
			CardNumber:     "4321-0987-6543-2109",
			Status:         models.ACTIVE_BANK_ACCOUNT,
		},
		{
			CustomerID:     7,
			AccountNumber:  "7890123456",
			Balance:        750.00,
			AccountType:    "savings",
			Cvc:            "456",
			ExpirationDate: "2026-02-28",
			CardNumber:     "3210-9876-5432-1098",
			Status:         models.ACTIVE_BANK_ACCOUNT,
		},
		{
			CustomerID:     8,
			AccountNumber:  "8901234567",
			Balance:        4000.00,
			AccountType:    "checking",
			Cvc:            "789",
			ExpirationDate: "2027-06-30",
			CardNumber:     "2109-8765-4321-0987",
			Status:         models.ACTIVE_BANK_ACCOUNT,
		},
		{
			CustomerID:     9,
			AccountNumber:  "9012345678",
			Balance:        1250.75,
			AccountType:    "savings",
			Cvc:            "321",
			ExpirationDate: "2025-12-01",
			CardNumber:     "1098-7654-3210-9876",
			Status:         models.ACTIVE_BANK_ACCOUNT,
		},
		{
			CustomerID:     10,
			AccountNumber:  "0123456789",
			Balance:        5000.00,
			AccountType:    "checking",
			Cvc:            "654",
			ExpirationDate: "2026-10-10",
			CardNumber:     "0987-6543-2109-8765",
			Status:         models.ACTIVE_BANK_ACCOUNT,
		},
	}

	gormdb := db.GetDbHandler().DB
	for _, item := range bankaccounts {
		gormdb.FirstOrCreate(&item, models.BankAccount{AccountNumber: item.AccountNumber})
	}
}
