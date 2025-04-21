package db

import (
	"testing"
	"yamanmnur/simple-dashboard/pkg/db"

	"github.com/DATA-DOG/go-sqlmock"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func TestInit(t *testing.T) {
	// Mock database URL for testing
	mockDb, _, _ := sqlmock.New()
	dialector := postgres.New(postgres.Config{
		Conn:       mockDb,
		DriverName: "postgres",
	})
	gormDb, _ := gorm.Open(dialector, &gorm.Config{})
	// Call the Init function
	dbpg := db.Init(gormDb)

	// Check if the returned db object is not nil
	if dbpg == nil {
		t.Fatal("Expected non-nil *gorm.DB, got nil")
	}

	// Check if the connection is valid
	sqlDB, err := dbpg.DB()
	if err != nil {
		t.Fatalf("Failed to get database connection: %v", err)
	}

	// Ping the database to ensure the connection is alive
	if err := sqlDB.Ping(); err != nil {
		t.Fatalf("Failed to ping database: %v", err)
	}
}
