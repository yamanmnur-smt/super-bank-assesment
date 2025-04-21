package repositories_test

import (
	"regexp"
	"testing"

	"yamanmnur/simple-dashboard/internal/dto/data"
	"yamanmnur/simple-dashboard/internal/repositories"
	"yamanmnur/simple-dashboard/pkg/db"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func TestGetTotalCard(t *testing.T) {
	mockDb, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer mockDb.Close()

	gormDB, err := gorm.Open(postgres.New(postgres.Config{
		Conn: mockDb,
	}), &gorm.Config{})
	assert.NoError(t, err)

	dbHandler := db.IDbHandler{
		DB: gormDB,
	}

	repo := &repositories.DashboardRepository{IDbHandler: &dbHandler}

	expectedResult := data.DashboardTotalCard{
		TotalCustomers: "100",
		TotalDeposits:  "50000",
		TotalBalance:   "200000",
	}

	mock.ExpectQuery(regexp.QuoteMeta(`
		SELECT
			(SELECT COUNT(id) FROM customers) AS total_customers,
			(SELECT SUM(amount) FROM term_deposits) AS total_deposits,
			(SELECT SUM(balance) FROM bank_accounts) AS total_balance
	`)).
		WillReturnRows(sqlmock.NewRows([]string{"total_customers", "total_deposits", "total_balance"}).
			AddRow("100", "50000", "200000"))

	result, err := repo.GetTotalCard()

	assert.NoError(t, err)
	assert.Equal(t, expectedResult, result)
}

func TestGetPieData(t *testing.T) {
	mockDb, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer mockDb.Close()

	gormDB, err := gorm.Open(postgres.New(postgres.Config{
		Conn: mockDb,
	}), &gorm.Config{})
	assert.NoError(t, err)

	dbHandler := db.IDbHandler{
		DB: gormDB,
	}

	repo := &repositories.DashboardRepository{IDbHandler: &dbHandler}

	expectedResult := []data.ChartData{
		{Label: "Type A", Value: "10000"},
		{Label: "Type B", Value: "20000"},
	}

	mock.ExpectQuery(regexp.QuoteMeta(`
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
	`)).
		WillReturnRows(
			sqlmock.NewRows([]string{"label", "value"}).
				AddRow("Type A", "10000").
				AddRow("Type B", "20000"))

	result, err := repo.GetPieData()

	assert.NoError(t, err)
	assert.Equal(t, expectedResult, result)
}

func TestGetMonthlyRegisteredCustomers(t *testing.T) {
	mockDb, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer mockDb.Close()

	gormDB, err := gorm.Open(postgres.New(postgres.Config{
		Conn: mockDb,
	}), &gorm.Config{})
	assert.NoError(t, err)

	dbHandler := db.IDbHandler{
		DB: gormDB,
	}

	repo := &repositories.DashboardRepository{IDbHandler: &dbHandler}

	expectedResult := []data.ChartData{
		{Label: "Jan", Value: "10"},
		{Label: "Feb", Value: "15"},
		{Label: "Mar", Value: "20"},
	}

	mock.ExpectQuery(regexp.QuoteMeta(`
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
		`)).
		WillReturnRows(
			sqlmock.NewRows([]string{"label", "value"}).
				AddRow("Jan", "10").
				AddRow("Feb", "15").
				AddRow("Mar", "20"))

	result, err := repo.GetMonthlyRegisteredCustomers()

	assert.NoError(t, err)
	assert.Equal(t, expectedResult, result)
}
