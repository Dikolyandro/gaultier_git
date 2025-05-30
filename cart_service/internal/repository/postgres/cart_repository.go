package postgres

import (
	"database/sql"
	"github.com/da-er-gaultier/cart_service/internal/domain"
)

type CartRepo struct {
	DB *sql.DB
}

func NewCartRepo(db *sql.DB) *CartRepo {
	return &CartRepo{DB: db}
}

func (r *CartRepo) Add(item *domain.CartItem) error {
	query := `INSERT INTO cart_items (user_id, product_id, quantity, created_at, updated_at)
			  VALUES ($1, $2, $3, NOW(), NOW()) RETURNING id`
	return r.DB.QueryRow(query, item.UserID, item.ProductID, item.Quantity).
		Scan(&item.ID)
}

func (r *CartRepo) GetByUser(userID int) ([]domain.CartItem, error) {
	query := `SELECT id, user_id, product_id, quantity, created_at, updated_at FROM cart_items WHERE user_id = $1`
	rows, err := r.DB.Query(query, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var items []domain.CartItem
	for rows.Next() {
		var i domain.CartItem
		if err := rows.Scan(&i.ID, &i.UserID, &i.ProductID, &i.Quantity, &i.CreatedAt, &i.UpdatedAt); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	return items, nil
}

func (r *CartRepo) Delete(itemID int) error {
	_, err := r.DB.Exec("DELETE FROM cart_items WHERE id = $1", itemID)
	return err
}
