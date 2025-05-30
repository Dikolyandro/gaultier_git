package domain

import "time"

type Delivery struct {
	ID        int       `json:"id"`
	OrderID   int       `json:"order_id"`
	Status    string    `json:"status"` // "packed", "shipped", "delivered"
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type DeliveryRepository interface {
	Create(delivery *Delivery) error
	UpdateStatus(id int, status string) error
	GetByID(id int) (*Delivery, error)
	GetByOrderID(orderID int) (*Delivery, error)
}

type DeliveryUsecase interface {
	Create(delivery *Delivery) error
	UpdateStatus(id int, status string) error
	GetByID(id int) (*Delivery, error)
	GetByOrderID(orderID int) (*Delivery, error)
}
