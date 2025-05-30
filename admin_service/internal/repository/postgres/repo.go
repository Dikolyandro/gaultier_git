package postgres

import (
	"github.com/da-er-gaultier/admin_service/internal/domain"
	"github.com/jmoiron/sqlx"
)

type postgresRepository struct {
	db *sqlx.DB
}

func New(db *sqlx.DB) domain.AdminRepository {
	return &postgresRepository{db: db}
}

// 🔽 USERS отключены, теперь работают через user_service

// 🔽 PRODUCTS
func (r *postgresRepository) GetAllProducts() ([]domain.Product, error) {
	var products []domain.Product
	err := r.db.Select(&products, "SELECT id, name, description, price FROM products")
	return products, err
}

func (r *postgresRepository) DeleteProductByID(id int) error {
	_, err := r.db.Exec("DELETE FROM products WHERE id = $1", id)
	return err
}

// 🔽 ORDERS
func (r *postgresRepository) GetAllOrders() ([]domain.Order, error) {
	var orders []domain.Order
	err := r.db.Select(&orders, "SELECT id, user_id, status FROM orders")
	return orders, err
}

func (r *postgresRepository) UpdateOrderStatus(orderID int, status string) error {
	_, err := r.db.Exec("UPDATE orders SET status = $1 WHERE id = $2", status, orderID)
	return err
}
