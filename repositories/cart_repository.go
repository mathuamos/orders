package repositories

import (
    "orders/entities"
    "github.com/jinzhu/gorm"
)

type CartRepository struct {
    db *gorm.DB
}



func NewCartRepository(db *gorm.DB) *CartRepository {
    return &CartRepository{db}
}



// GetAll retrieves all carts from the database.
func (r *CartRepository) GetAll() ([]entities.Cart, error) {
    var carts []entities.Cart
    if err := r.db.Find(&carts).Error; err != nil {
        return nil, err
    }
    return carts, nil
}




func (r *CartRepository) Create(cart *entities.Cart) error {
    return r.db.Create(cart).Error
}

func (r *CartRepository) GetByID(id uint) (*entities.Cart, error) {
    var cart entities.Cart
    if err := r.db.First(&cart, id).Error; err != nil {
        return nil, err
    }
    return &cart, nil
}


// repositories/cart_repository.go

func (r *CartRepository) Update(card *entities.Cart) error {
    if err := r.db.Save(card).Error; err != nil {
        return err
    }
    return nil
}



func (r *CartRepository) Delete(card *entities.Cart) error {
    if err := r.db.Delete(card).Error; err != nil {
        return err
    }
    return nil
}


// Implement other CRUD methods (Update, Delete, GetAll) similarly
