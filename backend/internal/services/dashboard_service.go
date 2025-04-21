package services

import (
	"yamanmnur/simple-dashboard/internal/dto/data"
	"yamanmnur/simple-dashboard/internal/repositories"
	"yamanmnur/simple-dashboard/pkg/util"
)

type IDashboardService interface {
	GetDashboard() (data.DashboardData, error)
}

type DashboardService struct {
	Repository repositories.IDashboardRepository
}

func (service *DashboardService) GetDashboard() (data.DashboardData, error) {
	total_card, err := service.Repository.GetTotalCard()
	if err != nil {
		return data.DashboardData{}, err
	}

	format_balance, _ := util.FormatIDRCurrency(total_card.TotalBalance)
	format_deposit, _ := util.FormatIDRCurrency(total_card.TotalDeposits)

	total_card.TotalBalance = format_balance
	total_card.TotalDeposits = format_deposit

	pie_data, err := service.Repository.GetPieData()
	for i, item := range pie_data {
		format_value, _ := util.FormatIDRCurrency(item.Value)

		pie_data[i].Value = format_value
	}
	if err != nil {
		return data.DashboardData{}, err
	}

	bar_data, err := service.Repository.GetMonthlyRegisteredCustomers()
	if err != nil {
		return data.DashboardData{}, err
	}

	return data.DashboardData{
		TotalCards: total_card,
		PieData:    pie_data,
		BarData:    bar_data,
	}, nil
}
