package seeder

import (
	"yamanmnur/simple-dashboard/internal/models"
	"yamanmnur/simple-dashboard/pkg/db"
)

func TermDepositTypesSeeder() {
	termDepositsTypes := []models.TermDepositsTypes{
		{
			Name:          "7 Days",
			InterestRate:  0.5,
			MinAmount:     0,
			MaxAmount:     1000000,
			TermDays:      7,
			EffectiveDate: "2023-01-01",
			Status:        "active",
		},
		{
			Name:          "14 Days",
			InterestRate:  0.6,
			MinAmount:     0,
			MaxAmount:     1000000,
			TermDays:      14,
			EffectiveDate: "2023-01-01",
			Status:        "active",
		},
		{
			Name:          "1 Month",
			InterestRate:  7.5,
			MinAmount:     500000,
			MaxAmount:     1000000000,
			TermDays:      30,
			EffectiveDate: "2023-01-01",
			Status:        "active",
		},
		{
			Name:          "3 Months",
			InterestRate:  7.5,
			MinAmount:     500000,
			MaxAmount:     1000000000,
			TermDays:      90,
			EffectiveDate: "2023-01-01",
			Status:        "active",
		},
		{
			Name:          "6 Months",
			InterestRate:  7.5,
			MinAmount:     500000,
			MaxAmount:     1000000000,
			TermDays:      180,
			EffectiveDate: "2023-01-01",
			Status:        "active",
		},
		{
			Name:          "9 Months",
			InterestRate:  7.5,
			MinAmount:     500000,
			MaxAmount:     1000000000,
			TermDays:      270,
			EffectiveDate: "2023-01-01",
			Status:        "active",
		},
		{
			Name:          "12 Months",
			InterestRate:  7.5,
			MinAmount:     500000,
			MaxAmount:     1000000000,
			TermDays:      365,
			EffectiveDate: "2023-01-01",
			Status:        "active",
		},
		{
			Name:          "1 Month",
			InterestRate:  6.5,
			MinAmount:     1000000001,
			MaxAmount:     0,
			TermDays:      30,
			EffectiveDate: "2023-01-01",
			Status:        "active",
		},
		{
			Name:          "3 Months",
			InterestRate:  6.5,
			MinAmount:     1000000001,
			MaxAmount:     0,
			TermDays:      90,
			EffectiveDate: "2023-01-01",
			Status:        "active",
		},
		{
			Name:          "6 Months",
			InterestRate:  6.5,
			MinAmount:     1000000001,
			MaxAmount:     0,
			TermDays:      180,
			EffectiveDate: "2023-01-01",
			Status:        "active",
		},
		{
			Name:          "9 Months",
			InterestRate:  6.5,
			MinAmount:     1000000001,
			MaxAmount:     0,
			TermDays:      270,
			EffectiveDate: "2023-01-01",
			Status:        "active",
		},
		{
			Name:          "12 Months",
			InterestRate:  6.5,
			MinAmount:     1000000001,
			MaxAmount:     0,
			TermDays:      365,
			EffectiveDate: "2023-01-01",
			Status:        "active",
		},
	}

	gormdb := db.GetDbHandler().DB

	var count int64
	gormdb.Model(&models.TermDepositsTypes{}).Count(&count)
	if count == 0 {
		for _, item := range termDepositsTypes {
			gormdb.Create(&item)
		}
	}
}
