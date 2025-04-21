package repositories_test

import (
	"fmt"
	"regexp"
	"testing"
	"yamanmnur/simple-dashboard/internal/dto/data"
	"yamanmnur/simple-dashboard/internal/models"
	"yamanmnur/simple-dashboard/internal/repositories"
	"yamanmnur/simple-dashboard/pkg/db"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func TestUserRepository_FindById(t *testing.T) {
	mockDb, mock, _ := sqlmock.New()
	dialector := postgres.New(postgres.Config{
		Conn:       mockDb,
		DriverName: "postgres",
	})
	gormDb, _ := gorm.Open(dialector, &gorm.Config{})

	dbHandler := db.IDbHandler{
		DB: gormDb,
	}

	repo := repositories.UserRepository{IDbHandler: &dbHandler}

	mockUserModel := models.User{
		Username: "yaman",
		Password: "jsdfjsldf",
		Name:     "name",
	}
	mockUserModel.ID = 1

	rows := sqlmock.NewRows([]string{
		"ID",
		"Username",
		"Password",
		"Name"}).AddRow(
		1,
		"yaman",
		"jsdfjsldf",
		"name")
	mock.ExpectQuery(`SELECT`).WillReturnRows(rows)
	users, err := repo.FindById(1)

	assert.NoError(t, err)
	assert.Equal(t, mockUserModel, users)
}

func TestUserRepository_FindByUsername(t *testing.T) {
	mockDb, mock, _ := sqlmock.New()
	dialector := postgres.New(postgres.Config{
		Conn:       mockDb,
		DriverName: "postgres",
	})
	gormDb, _ := gorm.Open(dialector, &gorm.Config{})

	dbHandler := db.IDbHandler{
		DB: gormDb,
	}

	repo := repositories.UserRepository{IDbHandler: &dbHandler}

	mockUserModel := models.User{
		Username: "yaman",
		Password: "jsdfjsldf",
		Name:     "name",
	}
	mockUserModel.ID = 1

	rows := sqlmock.NewRows([]string{
		"ID",
		"Username",
		"Password",
		"Name"}).AddRow(
		1,
		"yaman",
		"jsdfjsldf",
		"name")
	mock.ExpectQuery(`SELECT`).WillReturnRows(rows)
	users, err := repo.FindByUsername("yaman")

	assert.NoError(t, err)
	assert.Equal(t, mockUserModel, users)
}

func TestCreateUser(t *testing.T) {
	// Create a new SQL mock database
	database, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer database.Close()

	// Use GORM with the mock database
	gormDB, err := gorm.Open(postgres.New(postgres.Config{
		Conn: database,
	}), &gorm.Config{})
	assert.NoError(t, err)

	// Initialize repository
	dbHandler := db.IDbHandler{
		DB: gormDB,
	}
	t.Run("Create Success", func(t *testing.T) {
		userRepo := repositories.UserRepository{IDbHandler: &dbHandler}

		// Expected result
		expectedUser := models.User{
			Name:     "John Doe",
			Username: "johndoe",
			Password: "securepassword",
		}
		expectedUser.ID = 1 // Assuming the ID is auto-generated and returned

		// Sample input data
		inputUser := data.UserData{
			Name:     "John Doe",
			Username: "johndoe",
			Password: "securepassword",
		}

		// Mock SQL query expectations
		mock.ExpectQuery(regexp.QuoteMeta(
			`SELECT * FROM "users" WHERE username = $1 AND "users"."deleted_at" IS NULL ORDER BY "users"."id" LIMIT $2`)).
			WithArgs(inputUser.Username, 1).
			WillReturnRows(sqlmock.NewRows([]string{})) // No existing user found

		mock.ExpectBegin()
		mock.ExpectQuery(regexp.QuoteMeta(
			`INSERT INTO "users" ("created_at","updated_at","deleted_at","username","password","name") VALUES ($1,$2,$3,$4,$5,$6) RETURNING "id"`)).
			WithArgs(sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg(), inputUser.Username, inputUser.Password, inputUser.Name).
			WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(1))
		mock.ExpectCommit()

		// Call the repository function
		result, err := userRepo.Create(inputUser)

		// Assertions
		assert.NoError(t, err)
		assert.Equal(t, expectedUser.Name, result.Name)
		assert.Equal(t, expectedUser.Username, result.Username)
		assert.Equal(t, expectedUser.Password, result.Password)

		// Ensure all expectations were met
		assert.NoError(t, mock.ExpectationsWereMet())
	})

	t.Run("Create Error", func(t *testing.T) {

		userRepo := repositories.UserRepository{IDbHandler: &dbHandler}

		// Sample input data
		inputUser := data.UserData{
			Name:     "John Doe",
			Username: "johndoe",
			Password: "securepassword",
		}

		// Mock SQL query expectations
		mock.ExpectQuery(regexp.QuoteMeta(
			`SELECT * FROM "users" WHERE username = $1 AND "users"."deleted_at" IS NULL ORDER BY "users"."id" LIMIT $2`)).
			WithArgs(inputUser.Username, 1).
			WillReturnRows(sqlmock.NewRows([]string{})) // No existing user found

		mock.ExpectBegin()
		mock.ExpectQuery(regexp.QuoteMeta(
			`INSERT INTO "users" ("created_at","updated_at","deleted_at","username","password","name") VALUES ($1,$2,$3,$4,$5,$6) RETURNING "id"`)).
			WithArgs(sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg(), inputUser.Username, inputUser.Password, inputUser.Name).
			WillReturnError(fmt.Errorf("simulated db error"))
		mock.ExpectRollback()

		// Call the repository function
		result, err := userRepo.Create(inputUser)

		// Assertions
		assert.EqualError(t, err, "simulated db error")
		assert.Equal(t, models.User{}, result)
		assert.NoError(t, mock.ExpectationsWereMet())
	})

	t.Run("Create User Already Exists", func(t *testing.T) {

		userRepo := repositories.UserRepository{IDbHandler: &dbHandler}

		// Sample input data
		inputUser := data.UserData{
			Name:     "John Doe",
			Username: "johndoe",
			Password: "securepassword",
		}

		// Mock SQL query expectations
		mock.ExpectQuery(regexp.QuoteMeta(
			`SELECT * FROM "users" WHERE username = $1 AND "users"."deleted_at" IS NULL ORDER BY "users"."id" LIMIT $2`)).
			WithArgs(inputUser.Username, 1).
			// WillReturnRows(sqlmock.NewRows([]string{})) // No existing user found
			WillReturnRows(sqlmock.NewRows([]string{
				"id", "name", "username", "password",
			}).AddRow(1, "John Doe", "johndoe", "securepassword"))

		// Call the repository function
		result, err := userRepo.Create(inputUser)

		// Assertions
		assert.EqualError(t, err, "user exists")
		assert.Equal(t, models.User{}, result)
		assert.NoError(t, mock.ExpectationsWereMet())
	})
}

func TestUserRepository_Update(t *testing.T) {
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

	t.Run("Success Update User", func(t *testing.T) {
		repo := repositories.UserRepository{IDbHandler: &dbHandler}

		input := data.UserData{
			Id:       1,
			Name:     "Updated Name",
			Username: "Update Username",
			Password: "password",
		}

		mock.ExpectBegin()
		mock.ExpectExec(regexp.QuoteMeta(`UPDATE "users" SET`)).
			WithArgs(
				sqlmock.AnyArg(), // updated_at
				input.Username,
				input.Password,
				input.Name,
				input.Id,
			).
			WillReturnResult(sqlmock.NewResult(1, 1))
		mock.ExpectCommit()

		user, err := repo.Update(input)

		assert.NoError(t, err)
		assert.Equal(t, input.Name, user.Name)
		assert.Equal(t, input.Username, user.Username)
		assert.Equal(t, input.Password, user.Password)
		assert.NoError(t, mock.ExpectationsWereMet())
	})

	t.Run("DB Error", func(t *testing.T) {
		repo := repositories.UserRepository{IDbHandler: &dbHandler}

		input := data.UserData{
			Id:       1,
			Name:     "Updated Name",
			Username: "Update Username",
			Password: "password",
		}
		mock.ExpectBegin()
		mock.ExpectExec(regexp.QuoteMeta(`UPDATE "users" SET`)).
			WithArgs(
				sqlmock.AnyArg(), // updated_at
				input.Username,
				input.Password,
				input.Name,
				input.Id,
			).
			WillReturnError(fmt.Errorf("simulated db error"))
		mock.ExpectRollback()

		customer, err := repo.Update(input)

		assert.Error(t, err)
		assert.EqualError(t, err, "simulated db error")
		assert.Equal(t, models.User{}, customer)
		assert.NoError(t, mock.ExpectationsWereMet())
	})
}
