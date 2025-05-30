package domain

import "time"

type CartItem struct {
	ID        int       `json:"id"`
	UserID    int       `json:"user_id"`
	ProductID int       `json:"product_id"`
	Quantity  int       `json:"quantity"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type CartRepository interface {
	Add(item *CartItem) error
	GetByUser(userID int) ([]CartItem, error)
	Delete(itemID int) error
}

type CartUsecase interface {
	AddToCart(item *CartItem) error
	GetCart(userID int) ([]CartItem, error)
	RemoveFromCart(itemID int) error
}
