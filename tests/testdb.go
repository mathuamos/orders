package tests

import (
	"orders/entities"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

// SetupTestDB creates an in-memory SQLite database for testing.
func SetupTestDB() (*gorm.DB, func()) {
	db, err := gorm.Open("sqlite3", ":memory:")
	if err != nil {
		panic("Failed to connect to database: " + err.Error())
	}

	// Auto-migrate your database schema (create tables)
	db.AutoMigrate(&entities.Cart{})

	cleanup := func() {
		db.Close()
	}

	return db, cleanup
}
