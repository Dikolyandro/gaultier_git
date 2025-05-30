package domain

import "time"

type OrderItem struct {
	ProductID int `json:"product_id"`
	Quantity  int `json:"quantity"`
}

type Order struct {
	ID        int         `json:"id"`
	UserID    int         `json:"user_id"`
	Items     []OrderItem `json:"items"`
	Status    string      `json:"status"`
	CreatedAt time.Time   `json:"created_at"`
	UpdatedAt time.Time   `json:"updated_at"`
}

type OrderRepository interface {
	Create(order *Order) error
	GetByID(orderID int) (*Order, error)
	GetByUserID(userID int) ([]Order, error)
	Delete(orderID int) error
}

type OrderUsecase interface {
	Create(order *Order) error
	Get(orderID int) (*Order, error)
	GetByUser(userID int) ([]Order, error)
	Delete(orderID int) error
}
