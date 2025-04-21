package seeder

import (
	"yamanmnur/simple-dashboard/internal/models"
	"yamanmnur/simple-dashboard/pkg/db"
)

func PocketSeeder() {
	pockets := []models.Pocket{
		{
			CustomerID: 1,
			Name:       "Savings",
			Balance:    1000.50,
			Currency:   "USD",
		},
		{
			CustomerID: 2,
			Name:       "Travel Fund",
			Balance:    500.00,
			Currency:   "EUR",
		},
		{
			CustomerID: 3,
			Name:       "Emergency Fund",
			Balance:    2000.00,
			Currency:   "USD",
		},
		{
			CustomerID: 4,
			Name:       "Investment",
			Balance:    1500.75,
			Currency:   "GBP",
		},
		{
			CustomerID: 5,
			Name:       "Education",
			Balance:    3000.00,
			Currency:   "USD",
		},
		{
			CustomerID: 6,
			Name:       "Health Savings",
			Balance:    1200.00,
			Currency:   "USD",
		},
		{
			CustomerID: 7,
			Name:       "Vacation",
			Balance:    800.00,
			Currency:   "CAD",
		},
		{
			CustomerID: 8,
			Name:       "Car Fund",
			Balance:    2500.00,
			Currency:   "USD",
		},
		{
			CustomerID: 9,
			Name:       "Home Renovation",
			Balance:    4000.00,
			Currency:   "AUD",
		},
		{
			CustomerID: 10,
			Name:       "Wedding Fund",
			Balance:    3500.00,
			Currency:   "USD",
		},
	}

	gormdb := db.GetDbHandler().DB

	var count int64
	gormdb.Model(&models.Pocket{}).Count(&count)
	if count == 0 {
		for _, item := range pockets {
			gormdb.Create(&item)
		}
	}
}
