package entities

import "time"

// CartItem represents the "cart_items" table in the database.
type CartItem struct {
    ID             uint      `json:"id" gorm:"primaryKey"`
    CreatedAt      time.Time `json:"created_at"`
    Status         string    `json:"status"`
    UpdatedAt      time.Time `json:"updated_at"`
    CartID         uint      `json:"cart_id"`
    DiscountPercent int       `json:"discount_percent"`
    Price          float64   `json:"price"`
    ProductID      uint      `json:"product_id"`
    Quantity       int       `json:"quantity"`
}
