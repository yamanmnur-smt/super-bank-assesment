package repositories

import (
	"yamanmnur/simple-dashboard/internal/dto/data"
	"yamanmnur/simple-dashboard/pkg/db"
)

type IDashboardRepository interface {
	GetTotalCard() (data.DashboardTotalCard, error)
	GetPieData() ([]data.ChartData, error)
	GetMonthlyRegisteredCustomers() ([]data.ChartData, error)
}

type DashboardRepository struct {
	*db.IDbHandler
}

func (repository *DashboardRepository) GetTotalCard() (data.DashboardTotalCard, error) {
	var result data.DashboardTotalCard

	err := repository.DB.
		Raw(`
			SELECT
				(SELECT COUNT(id) FROM customers) AS total_customers,
				(SELECT SUM(amount) FROM term_deposits) AS total_deposits,
				(SELECT SUM(balance) FROM bank_accounts) AS total_balance
		`).Scan(&result).Error

	return result, err
}

func (repository *DashboardRepository) GetPieData() ([]data.ChartData, error) {
	var result []data.ChartData

	err := repository.DB.
		Raw(`
			SELECT 
				term_deposits_types.name AS label, 
				COALESCE(SUM(term_deposits.amount), 0) AS value
			FROM 
				term_deposits_types
			LEFT JOIN 
				term_deposits 
			ON 
				term_deposits.term_deposits_type_id = term_deposits_types.id
			GROUP BY 
				term_deposits_types.name
		`).Scan(&result).Error

	return result, err
}

func (repository *DashboardRepository) GetMonthlyRegisteredCustomers() ([]data.ChartData, error) {
	var result []data.ChartData

	err := repository.DB.
		Raw(`
			WITH months AS (
				SELECT TO_CHAR(date_trunc('month', d), 'Mon') AS label, EXTRACT(MONTH FROM d) AS month_order
				FROM generate_series(
					date_trunc('year', CURRENT_DATE),
					date_trunc('year', CURRENT_DATE) + interval '1 year - 1 day',
					interval '1 month'
				) AS d
			)
			SELECT 
				months.label, 
				COALESCE(COUNT(customers.id), 0) AS value
			FROM 
				months
			LEFT JOIN 
				customers 
			ON 
				EXTRACT(MONTH FROM customers.created_at) = months.month_order
				AND EXTRACT(YEAR FROM customers.created_at) = EXTRACT(YEAR FROM CURRENT_DATE)
			GROUP BY 
				months.label, months.month_order
			ORDER BY 
				months.month_order
		`).Scan(&result).Error

	return result, err
}
