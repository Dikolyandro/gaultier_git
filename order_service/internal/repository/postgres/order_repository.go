package postgres

import (
	"database/sql"
	"encoding/json"
	"errors"

	"github.com/da-er-gaultier/order_service/internal/domain"
)

type OrderRepo struct {
	DB *sql.DB
}

func NewOrderRepo(db *sql.DB) *OrderRepo {
	return &OrderRepo{DB: db}
}

func (r *OrderRepo) Create(order *domain.Order) error {
	itemsJSON, err := json.Marshal(order.Items)
	if err != nil {
		return err
	}

	query := `INSERT INTO orders (user_id, items, status, created_at, updated_at)
	          VALUES ($1, $2, $3, NOW(), NOW()) RETURNING id`
	return r.DB.QueryRow(query, order.UserID, itemsJSON, order.Status).
		Scan(&order.ID)
}

func (r *OrderRepo) GetByID(orderID int) (*domain.Order, error) {
	query := `SELECT id, user_id, items, status, created_at, updated_at FROM orders WHERE id = $1`
	row := r.DB.QueryRow(query, orderID)

	var o domain.Order
	var items []byte
	if err := row.Scan(&o.ID, &o.UserID, &items, &o.Status, &o.CreatedAt, &o.UpdatedAt); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}
	if err := json.Unmarshal(items, &o.Items); err != nil {
		return nil, err
	}
	return &o, nil
}

func (r *OrderRepo) GetByUserID(userID int) ([]domain.Order, error) {
	query := `SELECT id, user_id, items, status, created_at, updated_at FROM orders WHERE user_id = $1`
	rows, err := r.DB.Query(query, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var orders []domain.Order
	for rows.Next() {
		var o domain.Order
		var items []byte
		if err := rows.Scan(&o.ID, &o.UserID, &items, &o.Status, &o.CreatedAt, &o.UpdatedAt); err != nil {
			return nil, err
		}
		if err := json.Unmarshal(items, &o.Items); err != nil {
			return nil, err
		}
		orders = append(orders, o)
	}
	return orders, nil
}

func (r *OrderRepo) Delete(orderID int) error {
	_, err := r.DB.Exec("DELETE FROM orders WHERE id = $1", orderID)
	return err
}
