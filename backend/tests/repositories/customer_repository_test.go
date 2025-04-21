package repositories_test

import (
	"fmt"
	"regexp"
	"testing"
	"yamanmnur/simple-dashboard/internal/dto/data"
	"yamanmnur/simple-dashboard/internal/models"
	"yamanmnur/simple-dashboard/internal/repositories"
	"yamanmnur/simple-dashboard/pkg/db"
	pkg_requests "yamanmnur/simple-dashboard/pkg/requests"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var mockCustomerModel = models.Customer{
	Name:           "John Doe",
	Username:       "johndoe",
	Photo:          "photo.jpg",
	Email:          "jhon@mail.com",
	PhoneNumber:    "123456789",
	Address:        "123 Main St",
	Gender:         "male",
	AccountPurpose: "Personal",
	SourceOfIncome: "Salary",
	IncomePerMonth: "5000",
	Jobs:           "Software Engineer",
	Position:       "Senior",
	Industries:     "IT",
	CompanyName:    "Tech Co",
	AddressCompany: "456 Tech St",
	BankAccounts: []models.BankAccount{
		{
			CardNumber:     "0095-2340-2342-2342",
			AccountNumber:  "00002342345",
			Balance:        200000.0,
			AccountType:    "savings",
			Cvc:            "321",
			ExpirationDate: "2025-12-01",
			Status:         models.ACTIVE_BANK_ACCOUNT,
			TermDeposit: []models.TermDeposit{
				{
					Amount:                500000,
					InterestRate:          0.4,
					ExtensionInstructions: models.NO_ROLLOVER,
					StartDate:             "2024-05-02",
					MaturityDate:          "2024-05-02",
					Status:                models.ACTIVE,
					TermDepositsType: models.TermDepositsTypes{
						Name:          "7 days",
						InterestRate:  0.5,
						MinAmount:     500000,
						TermDays:      7,
						MaxAmount:     10000000,
						EffectiveDate: "2024-05-02",
						Status:        models.TERM_ACTIVE,
					},
				},
			},
		},
	},
	Pockets: []models.Pocket{
		{
			Name:     "Savings",
			Balance:  1000.0,
			Currency: "IDR",
		},
		{
			Name:     "Investments",
			Balance:  2000.0,
			Currency: "IDR",
		},
	},
}

func init() {
	mockCustomerModel.ID = 1
}

func TestCustomerRepository_FindById(t *testing.T) {
	mockDb, mock, _ := sqlmock.New()

	dialector := postgres.New(postgres.Config{
		Conn:       mockDb,
		DriverName: "postgres",
	})

	gormDb, _ := gorm.Open(dialector, &gorm.Config{})

	dbHandler := db.IDbHandler{
		DB: gormDb,
	}

	t.Run("Success Get Customer", func(t *testing.T) {
		repo := repositories.CustomerRepository{IDbHandler: &dbHandler}

		rows := sqlmock.NewRows([]string{
			"id", "name", "email", "phone_number", "address",
		}).AddRow(
			mockCustomerModel.ID,
			mockCustomerModel.Name,
			mockCustomerModel.Email,
			mockCustomerModel.PhoneNumber,
			mockCustomerModel.Address,
		)

		mock.ExpectQuery(`SELECT .* FROM "customers" WHERE customers.id = \$1 AND "customers"."deleted_at" IS NULL ORDER BY "customers"."id" LIMIT \$2`).
			WithArgs(1, 1).
			WillReturnRows(rows)

		bankAccountRows := sqlmock.NewRows([]string{
			"id", "customer_id", "card_number", "account_number", "balance", "account_type", "cvc", "expiration_date", "status",
		}).AddRow(
			1, mockCustomerModel.ID, "0095-2340-2342-2342", "00002342345", 200000.0, "savings", "321", "2025-12-01", models.ACTIVE_BANK_ACCOUNT,
		)

		mock.ExpectQuery(`SELECT .* FROM "bank_accounts" WHERE "bank_accounts"."customer_id" = \$1 AND "bank_accounts"."deleted_at" IS NULL`).
			WithArgs(mockCustomerModel.ID).
			WillReturnRows(bankAccountRows)

		termDepositRows := sqlmock.NewRows([]string{
			"id", "bank_account_id", "amount", "interest_rate", "maturity_date",
		}).AddRow(
			1, 1, 500000.0, 5.0, "2025-12-01",
		)

		mock.ExpectQuery(`SELECT .* FROM "term_deposits" WHERE "term_deposits"."bank_account_id" = \$1 AND "term_deposits"."deleted_at" IS NULL`).
			WithArgs(1).
			WillReturnRows(termDepositRows)

		pocketRows := sqlmock.NewRows([]string{
			"id", "customer_id", "name", "balance",
		}).AddRow(
			1, mockCustomerModel.ID, "Main Pocket", 150000.0,
		)

		mock.ExpectQuery(`SELECT .* FROM "pockets" WHERE "pockets"."customer_id" = \$1 AND "pockets"."deleted_at" IS NULL`).
			WithArgs(mockCustomerModel.ID).
			WillReturnRows(pocketRows)

		customer, err := repo.FindById(mockCustomerModel.ID)
		assert.NoError(t, err)
		assert.Equal(t, mockCustomerModel.Name, customer.Name)
		assert.Equal(t, mockCustomerModel.Email, customer.Email)
		assert.Equal(t, mockCustomerModel.PhoneNumber, customer.PhoneNumber)
		assert.Equal(t, mockCustomerModel.Address, customer.Address)
	})

	t.Run("Customer Not Found", func(t *testing.T) {
		repo := repositories.CustomerRepository{IDbHandler: &dbHandler}

		mock.ExpectQuery(`SELECT .* FROM "customers" WHERE customers.id = \$1 AND "customers"."deleted_at" IS NULL ORDER BY "customers"."id" LIMIT \$2`).
			WithArgs(2, 1).
			WillReturnError(gorm.ErrRecordNotFound)

		customer, err := repo.FindById(2)
		assert.Error(t, err)
		assert.Equal(t, gorm.ErrRecordNotFound, err)
		assert.Equal(t, models.Customer{}, customer)
	})

	t.Run("DB Error", func(t *testing.T) {
		repo := repositories.CustomerRepository{IDbHandler: &dbHandler}

		mock.ExpectQuery(`SELECT .* FROM "customers" WHERE customers.id = \$1 AND "customers"."deleted_at" IS NULL ORDER BY "customers"."id" LIMIT \$2`).
			WithArgs(mockCustomerModel.ID, 1).
			WillReturnError(gorm.ErrInvalidTransaction)

		customer, err := repo.FindById(mockCustomerModel.ID)
		assert.Error(t, err)
		assert.Equal(t, gorm.ErrInvalidTransaction, err)
		assert.Equal(t, models.Customer{}, customer)
	})
}

func TestCustomerRepository_Detail(t *testing.T) {
	mockDb, mock, _ := sqlmock.New()

	dialector := postgres.New(postgres.Config{
		Conn:       mockDb,
		DriverName: "postgres",
	})

	gormDb, _ := gorm.Open(dialector, &gorm.Config{})

	dbHandler := db.IDbHandler{
		DB: gormDb,
	}

	t.Run("Success Get Detail Customer", func(t *testing.T) {
		repo := repositories.CustomerRepository{IDbHandler: &dbHandler}

		rows := sqlmock.NewRows([]string{
			"id", "name", "email", "phone_number", "address",
		}).AddRow(
			mockCustomerModel.ID,
			mockCustomerModel.Name,
			mockCustomerModel.Email,
			mockCustomerModel.PhoneNumber,
			mockCustomerModel.Address,
		)

		mock.ExpectQuery(`SELECT .* FROM "customers" WHERE customers.id = \$1 AND "customers"."deleted_at" IS NULL ORDER BY "customers"."id" LIMIT \$2`).
			WithArgs(1, 1).
			WillReturnRows(rows)

		bankAccountRows := sqlmock.NewRows([]string{
			"id", "customer_id", "card_number", "account_number", "balance", "account_type", "cvc", "expiration_date", "status",
		}).AddRow(
			1, mockCustomerModel.ID, "0095-2340-2342-2342", "00002342345", 200000.0, "savings", "321", "2025-12-01", models.ACTIVE_BANK_ACCOUNT,
		)

		mock.ExpectQuery(`SELECT .* FROM "bank_accounts" WHERE "bank_accounts"."customer_id" = \$1 AND "bank_accounts"."deleted_at" IS NULL`).
			WithArgs(mockCustomerModel.ID).
			WillReturnRows(bankAccountRows)

		termDepositRows := sqlmock.NewRows([]string{
			"id", "bank_account_id", "amount", "interest_rate", "maturity_date",
		}).AddRow(
			1, 1, 500000.0, 5.0, "2025-12-01",
		)

		mock.ExpectQuery(`SELECT .* FROM "term_deposits" WHERE "term_deposits"."bank_account_id" = \$1 AND "term_deposits"."deleted_at" IS NULL`).
			WithArgs(1).
			WillReturnRows(termDepositRows)

		pocketRows := sqlmock.NewRows([]string{
			"id", "customer_id", "name", "balance",
		}).AddRow(
			1, mockCustomerModel.ID, "Main Pocket", 150000.0,
		)

		mock.ExpectQuery(`SELECT .* FROM "pockets" WHERE "pockets"."customer_id" = \$1 AND "pockets"."deleted_at" IS NULL`).
			WithArgs(mockCustomerModel.ID).
			WillReturnRows(pocketRows)

		var emptyCustomer models.Customer

		err := repo.Detail(mockCustomerModel.ID, &emptyCustomer)
		assert.NoError(t, err)
		assert.Equal(t, mockCustomerModel.Name, emptyCustomer.Name)
		assert.Equal(t, mockCustomerModel.Email, emptyCustomer.Email)
		assert.Equal(t, mockCustomerModel.PhoneNumber, emptyCustomer.PhoneNumber)
		assert.Equal(t, mockCustomerModel.Address, emptyCustomer.Address)
	})

	t.Run("Customer Detail Not Found", func(t *testing.T) {
		repo := repositories.CustomerRepository{IDbHandler: &dbHandler}

		mock.ExpectQuery(`SELECT .* FROM "customers" WHERE customers.id = \$1 AND "customers"."deleted_at" IS NULL ORDER BY "customers"."id" LIMIT \$2`).
			WithArgs(2, 1).
			WillReturnError(gorm.ErrRecordNotFound)
		var emptyCustomer models.Customer

		err := repo.Detail(2, &emptyCustomer)
		assert.Error(t, err)
		assert.Equal(t, gorm.ErrRecordNotFound, err)
		assert.Equal(t, models.Customer{}, emptyCustomer)
	})

	t.Run("DB Error", func(t *testing.T) {
		repo := repositories.CustomerRepository{IDbHandler: &dbHandler}

		mock.ExpectQuery(`SELECT .* FROM "customers" WHERE customers.id = \$1 AND "customers"."deleted_at" IS NULL ORDER BY "customers"."id" LIMIT \$2`).
			WithArgs(mockCustomerModel.ID, 1).
			WillReturnError(gorm.ErrInvalidTransaction)
		var emptyCustomer models.Customer

		err := repo.Detail(mockCustomerModel.ID, &emptyCustomer)
		assert.Error(t, err)
		assert.Equal(t, gorm.ErrInvalidTransaction, err)
		assert.Equal(t, models.Customer{}, emptyCustomer)
	})
}

func TestCustomerRepository_Create(t *testing.T) {
	// Create a new SQL mock database
	mockDb, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer mockDb.Close()

	// Use GORM with the mock database
	gormDB, err := gorm.Open(postgres.New(postgres.Config{
		Conn: mockDb,
	}), &gorm.Config{})
	assert.NoError(t, err)

	// Initialize repository
	dbHandler := db.IDbHandler{
		DB: gormDB,
	}

	t.Run("Success Create Customer", func(t *testing.T) {
		repo := repositories.CustomerRepository{IDbHandler: &dbHandler}
		// Sample input data
		input := &models.Customer{
			Photo:          "photo.jpg",
			Name:           "John Doe",
			Email:          "jhon@mail.com",
			PhoneNumber:    "123456789",
			Address:        "123 Main St",
			Username:       "johndoe",
			Gender:         "male",
			AccountPurpose: "Personal",
			SourceOfIncome: "Salary",
			IncomePerMonth: "5000",
			Jobs:           "Software Engineer",
			Position:       "Senior",
			Industries:     "IT",
			CompanyName:    "Tech Co",
			AddressCompany: "456 Tech St",
		}

		// Mock SQL query expectations
		mock.ExpectBegin()
		mock.ExpectQuery(regexp.QuoteMeta(`INSERT INTO "customers"`)).
			WithArgs(
				sqlmock.AnyArg(), sqlmock.AnyArg(), nil,
				input.Photo,
				input.Name,
				input.Email,
				input.PhoneNumber,
				input.Address,
				input.Username,
				input.Gender,
				input.AccountPurpose,
				input.SourceOfIncome,
				input.IncomePerMonth,
				input.Jobs,
				input.Position,
				input.Industries,
				input.CompanyName,
				input.AddressCompany,
			).
			WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(1))

		mock.ExpectCommit()

		// Call the repository function
		err := repo.Create(input)

		// Assertions
		assert.NoError(t, err)
		assert.NoError(t, mock.ExpectationsWereMet())
	})

	t.Run("Fail Create Customer - db Error", func(t *testing.T) {
		repo := repositories.CustomerRepository{IDbHandler: &dbHandler}

		input := &models.Customer{
			Photo:          "photo.jpg",
			Name:           "John Doe",
			Email:          "jhon@mail.com",
			PhoneNumber:    "123456789",
			Address:        "123 Main St",
			Username:       "johndoe",
			Gender:         "male",
			AccountPurpose: "Personal",
			SourceOfIncome: "Salary",
			IncomePerMonth: "5000",
			Jobs:           "Software Engineer",
			Position:       "Senior",
			Industries:     "IT",
			CompanyName:    "Tech Co",
			AddressCompany: "456 Tech St",
		}

		mock.ExpectBegin()
		mock.ExpectQuery(regexp.QuoteMeta(`INSERT INTO "customers"`)).
			WithArgs(
				sqlmock.AnyArg(), // created_at
				sqlmock.AnyArg(), // updated_at
				nil,
				input.Photo,
				input.Name,
				input.Email,
				input.PhoneNumber,
				input.Address,
				input.Username,
				input.Gender,
				input.AccountPurpose,
				input.SourceOfIncome,
				input.IncomePerMonth,
				input.Jobs,
				input.Position,
				input.Industries,
				input.CompanyName,
				input.AddressCompany,
			).
			WillReturnError(fmt.Errorf("simulated db error"))
		mock.ExpectRollback()

		// Call the repo
		err = repo.Create(input)

		assert.Error(t, err)
		assert.EqualError(t, err, "simulated db error")
		assert.NoError(t, mock.ExpectationsWereMet())
	})
}
func TestCustomerRepository_Update(t *testing.T) {
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

	t.Run("Success Update Customer", func(t *testing.T) {
		repo := repositories.CustomerRepository{IDbHandler: &dbHandler}

		input := &data.CustomerData{
			Id:          1,
			Name:        "Updated Name",
			Email:       "updated@mail.com",
			PhoneNumber: "987654321",
			Address:     "Updated Address",
		}

		mock.ExpectBegin()
		mock.ExpectExec(regexp.QuoteMeta(`UPDATE "customers" SET`)).
			WithArgs(
				sqlmock.AnyArg(),
				input.Name,
				input.Email,
				input.PhoneNumber,
				input.Address,
				input.Id,
			).
			WillReturnResult(sqlmock.NewResult(1, 1))
		mock.ExpectCommit()

		customer, err := repo.Update(input)

		assert.NoError(t, err)
		assert.Equal(t, input.Name, customer.Name)
		assert.Equal(t, input.Email, customer.Email)
		assert.Equal(t, input.PhoneNumber, customer.PhoneNumber)
		assert.Equal(t, input.Address, customer.Address)
		assert.NoError(t, mock.ExpectationsWereMet())
	})

	t.Run("Customer Not Found", func(t *testing.T) {
		repo := repositories.CustomerRepository{IDbHandler: &dbHandler}

		input := &data.CustomerData{
			Id:          2,
			Name:        "Nonexistent Name",
			Email:       "nonexistent@mail.com",
			PhoneNumber: "000000000",
			Address:     "Nonexistent Address",
		}

		mock.ExpectBegin()
		mock.ExpectExec(regexp.QuoteMeta(`UPDATE "customers" SET`)).
			WithArgs(
				sqlmock.AnyArg(), // updated_at
				input.Name,
				input.Email,
				input.PhoneNumber,
				input.Address,
				input.Id,
			).
			WillReturnResult(sqlmock.NewResult(0, 0))
		mock.ExpectCommit()

		customer, err := repo.Update(input)

		assert.Error(t, err)
		assert.Equal(t, gorm.ErrRecordNotFound, err)
		assert.Equal(t, models.Customer{}, customer)
		assert.NoError(t, mock.ExpectationsWereMet())
	})

	t.Run("DB Error", func(t *testing.T) {
		repo := repositories.CustomerRepository{IDbHandler: &dbHandler}

		input := &data.CustomerData{
			Id:          1,
			Name:        "Error Name",
			Email:       "error@mail.com",
			PhoneNumber: "111111111",
			Address:     "Error Address",
		}

		mock.ExpectBegin()
		mock.ExpectExec(regexp.QuoteMeta(`UPDATE "customers" SET`)).
			WithArgs(
				sqlmock.AnyArg(), // updated_at
				input.Name,
				input.Email,
				input.PhoneNumber,
				input.Address,
				input.Id,
			).
			WillReturnError(fmt.Errorf("simulated db error"))
		mock.ExpectRollback()

		customer, err := repo.Update(input)

		assert.Error(t, err)
		assert.EqualError(t, err, "simulated db error")
		assert.Equal(t, models.Customer{}, customer)
		assert.NoError(t, mock.ExpectationsWereMet())
	})
}
func TestCustomerRepository_UpdatePatch(t *testing.T) {
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

	t.Run("Success UpdatePatch Customer", func(t *testing.T) {
		repo := repositories.CustomerRepository{IDbHandler: &dbHandler}

		input := &data.CustomerData{
			Id:          1,
			Name:        "Updated Name",
			Email:       "updated@mail.com",
			PhoneNumber: "987654321",
			Address:     "Updated Address",
			Photo:       "updated_photo.jpg",
		}

		mock.ExpectBegin()
		mock.ExpectExec(regexp.QuoteMeta(`UPDATE "customers" SET`)).
			WithArgs(
				sqlmock.AnyArg(), // updated_at
				input.Photo,
				input.Name,
				input.Email,
				input.PhoneNumber,
				input.Address,
				input.Id,
			).
			WillReturnResult(sqlmock.NewResult(1, 1))
		mock.ExpectCommit()

		customer, err := repo.UpdatePatch(input)

		assert.NoError(t, err)
		assert.Equal(t, input.Name, customer.Name)
		assert.Equal(t, input.Email, customer.Email)
		assert.Equal(t, input.PhoneNumber, customer.PhoneNumber)
		assert.Equal(t, input.Address, customer.Address)
		assert.Equal(t, input.Photo, customer.Photo)
		assert.NoError(t, mock.ExpectationsWereMet())
	})

	t.Run("Customer Not Found", func(t *testing.T) {
		repo := repositories.CustomerRepository{IDbHandler: &dbHandler}

		input := &data.CustomerData{
			Id:          2,
			Name:        "Nonexistent Name",
			Email:       "nonexistent@mail.com",
			PhoneNumber: "000000000",
			Address:     "Nonexistent Address",
		}

		mock.ExpectBegin()
		mock.ExpectExec(regexp.QuoteMeta(`UPDATE "customers" SET`)).
			WithArgs(
				sqlmock.AnyArg(), // updated_at
				input.Name,
				input.Email,
				input.PhoneNumber,
				input.Address,
				input.Id,
			).
			WillReturnResult(sqlmock.NewResult(0, 0))
		mock.ExpectCommit()

		customer, err := repo.UpdatePatch(input)

		assert.Error(t, err)
		assert.Equal(t, gorm.ErrRecordNotFound, err)
		assert.Equal(t, models.Customer{}, customer)
		assert.NoError(t, mock.ExpectationsWereMet())
	})

	t.Run("DB Error", func(t *testing.T) {
		repo := repositories.CustomerRepository{IDbHandler: &dbHandler}

		input := &data.CustomerData{
			Id:          1,
			Name:        "Error Name",
			Email:       "error@mail.com",
			PhoneNumber: "111111111",
			Address:     "Error Address",
			Photo:       "error_photo.jpg",
		}

		mock.ExpectBegin()
		mock.ExpectExec(regexp.QuoteMeta(`UPDATE "customers" SET`)).
			WithArgs(
				sqlmock.AnyArg(), // updated_at
				input.Photo,
				input.Name,
				input.Email,
				input.PhoneNumber,
				input.Address,
				input.Id,
			).
			WillReturnError(fmt.Errorf("simulated db error"))
		mock.ExpectRollback()

		customer, err := repo.UpdatePatch(input)

		assert.Error(t, err)
		assert.EqualError(t, err, "simulated db error")
		assert.Equal(t, models.Customer{}, customer)
		assert.NoError(t, mock.ExpectationsWereMet())
	})

	t.Run("Invalid Input - Missing ID", func(t *testing.T) {
		repo := repositories.CustomerRepository{IDbHandler: &dbHandler}

		input := &data.CustomerData{
			Id:          0,
			Name:        "Invalid Name",
			Email:       "invalid@mail.com",
			PhoneNumber: "000000000",
			Address:     "Invalid Address",
		}

		customer, err := repo.UpdatePatch(input)

		assert.Error(t, err)
		assert.Equal(t, gorm.ErrRecordNotFound, err)
		assert.Equal(t, models.Customer{}, customer)
	})
}
func TestCustomerRepository_Delete(t *testing.T) {
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

	t.Run("Success Delete Customer", func(t *testing.T) {
		repo := repositories.CustomerRepository{IDbHandler: &dbHandler}

		mock.ExpectBegin()
		mock.ExpectExec(regexp.QuoteMeta(`UPDATE "customers" SET`)).
			WithArgs(sqlmock.AnyArg(), 1).
			WillReturnResult(sqlmock.NewResult(1, 1))
		mock.ExpectCommit()

		err := repo.Delete(1)

		assert.NoError(t, err)
		assert.NoError(t, mock.ExpectationsWereMet())
	})

	t.Run("Customer Not Found", func(t *testing.T) {
		repo := repositories.CustomerRepository{IDbHandler: &dbHandler}

		mock.ExpectBegin()
		mock.ExpectExec(regexp.QuoteMeta(`UPDATE "customers" SET`)).
			WithArgs(sqlmock.AnyArg(), 2).
			WillReturnResult(sqlmock.NewResult(0, 0))
		mock.ExpectCommit()

		err := repo.Delete(2)

		assert.Error(t, err)
		assert.Equal(t, gorm.ErrRecordNotFound, err)
		assert.NoError(t, mock.ExpectationsWereMet())
	})

}
func TestCustomerRepository_GetCustomersWithPagination(t *testing.T) {
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

	t.Run("Success Get Customers With Pagination", func(t *testing.T) {
		repo := repositories.CustomerRepository{IDbHandler: &dbHandler}

		pageRequest := pkg_requests.PageRequest{
			PageNumber:    0,
			PageSize:      10,
			Search:        "",
			SortBy:        "name",
			SortDirection: "asc",
		}
		// Simulasi query COUNT
		mock.ExpectQuery(regexp.QuoteMeta(
			`SELECT count(*) FROM "customers" LEFT JOIN LATERAL (
			select 
				bank_accounts.id, 
				bank_accounts.account_number, 
				bank_accounts.customer_id 
			from bank_accounts 
			where bank_accounts.deleted_at is null
			and bank_accounts.customer_id = customers.id
			order by created_at desc  
			limit 1
		) b ON true`)).
			WillReturnRows(sqlmock.NewRows([]string{"count"}).AddRow(1))

		// Simulasi query data customer
		mock.ExpectQuery(regexp.QuoteMeta(`
				SELECT customers.id, 
					customers.name, 
					customers.email, 
					customers.phone_number, 
					customers.address, 
					TO_CHAR(customers.created_at, 'DD Mon YYYY HH24:MI:SS') as created_at, 
					b.account_number 
				FROM "customers" 
				LEFT JOIN LATERAL (
					select 
						bank_accounts.id, 
						bank_accounts.account_number, 
						bank_accounts.customer_id 
					from bank_accounts 
					where bank_accounts.deleted_at is null
					and bank_accounts.customer_id = customers.id
					order by created_at desc  
					limit 1
				) b ON true 
				WHERE "customers"."deleted_at" IS NULL 
				ORDER BY name asc LIMIT $1
			`)).
			WithArgs(10).
			WillReturnRows(sqlmock.NewRows([]string{
				"id", "name", "email", "phone_number", "address", "created_at", "account_number",
			}).AddRow(1, "John Doe", "john@mail.com", "08123456789", "Jl. Sudirman", "20 Apr 2025 20:10:00", "123456"))

		result, err := repo.GetCustomersWithPagination(pageRequest)

		assert.NoError(t, err)
		assert.Len(t, result.Data, 1)
		assert.Equal(t, "John Doe", result.Data[0].Name)
		assert.NoError(t, mock.ExpectationsWereMet())
	})

	t.Run("Search Customers With Pagination", func(t *testing.T) {
		repo := repositories.CustomerRepository{IDbHandler: &dbHandler}

		pageRequest := pkg_requests.PageRequest{
			PageNumber:    0,
			PageSize:      10,
			Search:        "John",
			SortBy:        "name",
			SortDirection: "asc",
		}

		mock.ExpectQuery(regexp.QuoteMeta(`SELECT count(*) FROM "customers" LEFT JOIN LATERAL ( select bank_accounts.id, bank_accounts.account_number, bank_accounts.customer_id from bank_accounts where bank_accounts.deleted_at is null and bank_accounts.customer_id = customers.id order by created_at desc limit 1 ) b ON true WHERE CONCAT_WS(' ', name, email, phone_number, b.account_number, TO_CHAR(customers.created_at, 'DD Mon YYYY HH24:MI:SS'), address) LIKE $1 AND "customers"."deleted_at" IS NULL`)).
			WithArgs("%John%").
			WillReturnRows(sqlmock.NewRows([]string{"count"}).AddRow(1))

		mock.ExpectQuery(regexp.QuoteMeta(`SELECT customers.id, customers.name, customers.email, customers.phone_number, customers.address, TO_CHAR(customers.created_at, 'DD Mon YYYY HH24:MI:SS') as created_at, b.account_number FROM "customers" LEFT JOIN LATERAL ( select bank_accounts.id, bank_accounts.account_number, bank_accounts.customer_id from bank_accounts where bank_accounts.deleted_at is null and bank_accounts.customer_id = customers.id order by created_at desc limit 1 ) b ON true WHERE CONCAT_WS(' ', name, email, phone_number, b.account_number, TO_CHAR(customers.created_at, 'DD Mon YYYY HH24:MI:SS'), address) LIKE $1 AND "customers"."deleted_at" IS NULL ORDER BY name asc LIMIT $2`)).
			WithArgs("%John%", 10).
			WillReturnRows(sqlmock.NewRows([]string{
				"id", "name", "email", "phone_number", "address", "created_at", "account_number",
			}).AddRow(
				1, "John Doe", "johndoe@mail.com", "123456789", "123 Main St", "01 Jan 2023 12:00:00", "00002342345",
			))

		result, err := repo.GetCustomersWithPagination(pageRequest)

		assert.NoError(t, err)
		assert.Equal(t, 1, len(result.Data))
		assert.Equal(t, int64(1), result.PageData.TotalRows)
		assert.Equal(t, 0, result.PageData.Page)
		assert.Equal(t, 10, result.PageData.Limit)
		assert.Equal(t, 1, result.PageData.TotalPages)
		assert.NoError(t, mock.ExpectationsWereMet())
	})

	t.Run("DB Error Count", func(t *testing.T) {
		repo := repositories.CustomerRepository{IDbHandler: &dbHandler}

		pageRequest := pkg_requests.PageRequest{
			PageNumber:    1,
			PageSize:      10,
			Search:        "",
			SortBy:        "name",
			SortDirection: "asc",
		}

		mock.ExpectQuery(regexp.QuoteMeta(
			`SELECT count(*) FROM "customers" LEFT JOIN LATERAL (
				select 
					bank_accounts.id, 
					bank_accounts.account_number, 
					bank_accounts.customer_id 
				from bank_accounts 
				where bank_accounts.deleted_at is null
				and bank_accounts.customer_id = customers.id
				order by created_at desc  
				limit 1
			) b ON true`)).
			WillReturnError(fmt.Errorf("simulated db error"))

		result, err := repo.GetCustomersWithPagination(pageRequest)

		assert.Error(t, err)
		assert.EqualError(t, err, "simulated db error")
		assert.Equal(t, 0, len(result.Data))
		assert.NoError(t, mock.ExpectationsWereMet())
	})

	t.Run("DB Error Limit", func(t *testing.T) {
		repo := repositories.CustomerRepository{IDbHandler: &dbHandler}

		pageRequest := pkg_requests.PageRequest{
			PageNumber:    0,
			PageSize:      10,
			Search:        "",
			SortBy:        "name",
			SortDirection: "asc",
		}

		mock.ExpectQuery(regexp.QuoteMeta(
			`SELECT count(*) FROM "customers" LEFT JOIN LATERAL (
				select 
					bank_accounts.id, 
					bank_accounts.account_number, 
					bank_accounts.customer_id 
				from bank_accounts 
				where bank_accounts.deleted_at is null
				and bank_accounts.customer_id = customers.id
				order by created_at desc  
				limit 1
			) b ON true`)).
			WillReturnRows(sqlmock.NewRows([]string{"count"}).AddRow(1))

		mock.ExpectQuery(regexp.QuoteMeta(`
				SELECT customers.id, 
					customers.name, 
					customers.email, 
					customers.phone_number, 
					customers.address, 
					TO_CHAR(customers.created_at, 'DD Mon YYYY HH24:MI:SS') as created_at, 
					b.account_number 
				FROM "customers" 
				LEFT JOIN LATERAL (
					select 
						bank_accounts.id, 
						bank_accounts.account_number, 
						bank_accounts.customer_id 
					from bank_accounts 
					where bank_accounts.deleted_at is null
					and bank_accounts.customer_id = customers.id
					order by created_at desc  
					limit 1
				) b ON true 
				WHERE "customers"."deleted_at" IS NULL 
				ORDER BY name asc LIMIT $1
			`)).
			WithArgs(10).
			WillReturnError(fmt.Errorf("simulated db error"))

		result, err := repo.GetCustomersWithPagination(pageRequest)

		assert.Error(t, err)
		assert.EqualError(t, err, "simulated db error")
		assert.Equal(t, 0, len(result.Data))
		assert.NoError(t, mock.ExpectationsWereMet())
	})
}
