package tests

import (
	"testing"

	"github.com/jinzhu/gorm"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"orders/entities"
	"orders/repositories"
)

func TestCartRepository(t *testing.T) {
	db, cleanup := SetupTestDB()
	defer cleanup()

	t.Run("TestCreateAndGetByID", func(t *testing.T) {
		testCreateAndGetByID(t, db)
	})

	t.Run("TestUpdateAndDelete", func(t *testing.T) {
		testUpdateAndDelete(t, db)
	})

	t.Run("TestGetAll", func(t *testing.T) {
		testGetAll(t, db)
	})
}

func testCreateAndGetByID(t *testing.T, db *gorm.DB) {
	cartRepo := repositories.NewCartRepository(db)

	// Test Create method
	cart := &entities.Cart{
		Status:      "Active",
		ItemsCount:  5,
		UserID:      1,
		TotalAmount: 100.0,
	}
	err := cartRepo.Create(cart)
	require.NoError(t, err)

	// Test GetByID method
	retrievedCart, err := cartRepo.GetByID(cart.ID)
	require.NoError(t, err)
	assert.Equal(t, "Active", retrievedCart.Status)
}

func testUpdateAndDelete(t *testing.T, db *gorm.DB) {
	cartRepo := repositories.NewCartRepository(db)

	// Create a cart
	cart := &entities.Cart{
		Status:      "Active",
		ItemsCount:  5,
		UserID:      1,
		TotalAmount: 100.0,
	}
	err := cartRepo.Create(cart)
	require.NoError(t, err)

	// Test Update method
	cart.Status = "Updated"
	err = cartRepo.Update(cart)
	require.NoError(t, err)

	// Test Delete method
	err = cartRepo.Delete(cart)
	require.NoError(t, err)

	// Verify that the cart is deleted
	_, err = cartRepo.GetByID(cart.ID)
	assert.Error(t, err)
}


func testGetAll(t *testing.T, db *gorm.DB) {
    cartRepo := repositories.NewCartRepository(db)

    // Clear the database table
    db.Delete(&entities.Cart{})

    // Create multiple carts
    for i := 0; i < 3; i++ {
        cart := &entities.Cart{
            Status:      "Active",
            ItemsCount:  5,
            UserID:      1,
            TotalAmount: 100.0,
        }
        err := cartRepo.Create(cart)
        require.NoError(t, err)
    }

    // Test GetAll method
    carts, err := cartRepo.GetAll()
    require.NoError(t, err)
    assert.Len(t, carts, 3)
}
