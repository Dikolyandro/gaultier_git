package postgres

import (
	"database/sql"
	"errors"
	"time"

	"github.com/da-er-gaultier/delivery_service/internal/domain"
)

type DeliveryRepo struct {
	DB *sql.DB
}

func NewDeliveryRepo(db *sql.DB) *DeliveryRepo {
	return &DeliveryRepo{DB: db}
}

func (r *DeliveryRepo) Create(d *domain.Delivery) error {
	query := `INSERT INTO deliveries (order_id, status, created_at, updated_at)
	          VALUES ($1, $2, NOW(), NOW()) RETURNING id`
	return r.DB.QueryRow(query, d.OrderID, d.Status).
		Scan(&d.ID)
}

func (r *DeliveryRepo) UpdateStatus(id int, status string) error {
	query := `UPDATE deliveries SET status = $1, updated_at = $2 WHERE id = $3`
	_, err := r.DB.Exec(query, status, time.Now(), id)
	return err
}

func (r *DeliveryRepo) GetByID(id int) (*domain.Delivery, error) {
	query := `SELECT id, order_id, status, created_at, updated_at FROM deliveries WHERE id = $1`
	row := r.DB.QueryRow(query, id)

	var d domain.Delivery
	if err := row.Scan(&d.ID, &d.OrderID, &d.Status, &d.CreatedAt, &d.UpdatedAt); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}
	return &d, nil
}

func (r *DeliveryRepo) GetByOrderID(orderID int) (*domain.Delivery, error) {
	query := `SELECT id, order_id, status, created_at, updated_at FROM deliveries WHERE order_id = $1`
	row := r.DB.QueryRow(query, orderID)

	var d domain.Delivery
	if err := row.Scan(&d.ID, &d.OrderID, &d.Status, &d.CreatedAt, &d.UpdatedAt); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}
	return &d, nil
}
