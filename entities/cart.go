package entities

type Cart struct {
    ID          uint    `json:"id" gorm:"primaryKey"`
    Status      string  `json:"status"`
    ItemsCount  int     `json:"items_count"`
    UserID      uint    `json:"user_id"`
    CreatedAt   string  `json:"created_at"` // Change the data type to string
    UpdatedAt   string  `json:"updated_at"` // Change the data type to string
    TotalAmount float64 `json:"total_amount"`
}

// TableName returns the table name for the Cart model.
func (Cart) TableName() string {
    return "cart" // Set the table name to "cart"
}
