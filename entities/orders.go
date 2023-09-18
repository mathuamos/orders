package entities

import "time"

// Order represents the "orders" table in the database.
type Order struct {
    ID            uint      `json:"id" gorm:"primaryKey"`
    CreatedAt     time.Time `json:"created_at"`
    Status        string    `json:"status"`
    UpdatedAt     time.Time `json:"updated_at"`
    CartID        uint      `json:"cart_id"`
    Comment       string    `json:"comment"`
    ItemsCount    int       `json:"items_count"`
    OrderStatus   string    `json:"order_status"`
    PaymentMethod string    `json:"payment_method"`
    TotalAmount   float64   `json:"total_amount"`
    UserID        uint      `json:"user_id"`
}
