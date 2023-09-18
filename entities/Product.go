package entities

import "time"

type Product struct {
    ID             uint `gorm:"primaryKey"`
    CreatedAt      time.Time
    Status         string
    UpdatedAt      time.Time
    CategoryID     uint
    Description    string
    DiscountPercent int
    Name           string
    Price          float64
    Priority       int
    StockQuantity  int
}
