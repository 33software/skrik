package tests

import (
	"skrik/internal/entities"
	"testing"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)
var TestDb *gorm.DB

func TestMain (m *testing.M) {
	dsn := "host=localhost port=5433 user=test password=test dbname=test_db sslmode=disable"
	var err error
    TestDb, err= gorm.Open(postgres.Open(dsn), &gorm.Config{})
    if err != nil {
        panic("failed to connect to test database")
    }
	TestDb.AutoMigrate(&entities.User{})

	m.Run()
}