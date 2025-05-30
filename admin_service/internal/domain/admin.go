package domain

import "time"

type User struct {
	ID        int       `json:"id"`
	Email     string    `json:"email"`
	Role      string    `json:"role"` // "user" or "admin"
	CreatedAt time.Time `json:"created_at"`
}

type Product struct {
	ID          int       `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Price       float64   `json:"price"`
	CreatedAt   time.Time `json:"created_at"`
}

type Order struct {
	ID        int       `json:"id"`
	UserID    int       `json:"user_id"`
	Status    string    `json:"status"`
	CreatedAt time.Time `json:"created_at"`
}

type AdminUsecase interface {
	ListUsers() ([]User, error)
	DeleteUser(userID int) error

	ListProducts() ([]Product, error)
	DeleteProduct(productID int) error

	ListOrders() ([]Order, error)
	UpdateOrderStatus(orderID int, status string) error
}

type AdminRepository interface {
	// Product
	GetAllProducts() ([]Product, error)
	DeleteProductByID(id int) error

	// Order
	GetAllOrders() ([]Order, error)
	UpdateOrderStatus(orderID int, status string) error
}
