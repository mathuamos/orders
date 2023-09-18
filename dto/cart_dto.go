package dto

import (
    "github.com/go-playground/validator/v10"
)

var validate = validator.New()

// CartDTO represents the data required to create a card.
type CartDTO struct {
    Status      string  `json:"status" validate:"required"`
    ItemsCount  int     `json:"items_count" validate:"required,min=1"`
    UserID      uint    `json:"user_id" validate:"required"`
    TotalAmount float64 `json:"total_amount" validate:"required,min=0"`
}

// Validate validates the CartDTO and returns any validation errors.
func (cartDTO *CartDTO) Validate() error {
    return validate.Struct(cartDTO)
}
