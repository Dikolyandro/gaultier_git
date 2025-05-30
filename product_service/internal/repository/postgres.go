package repository

import (
	"github.com/jmoiron/sqlx"
	"product_service/internal/domain"
)

type productRepository struct {
	db *sqlx.DB
}

func NewProductRepository(db *sqlx.DB) domain.ProductRepository {
	return &productRepository{db: db}
}

func (r *productRepository) Create(p *domain.Product) error {
	query := `
		INSERT INTO products (name, description, price, category, stock)
		VALUES ($1, $2, $3, $4, $5)
		RETURNING id, created_at, updated_at`
	return r.db.QueryRowx(query, p.Name, p.Description, p.Price, p.Category, p.Stock).
		Scan(&p.ID, &p.CreatedAt, &p.UpdatedAt)
}

func (r *productRepository) GetByID(id int64) (*domain.Product, error) {
	var p domain.Product
	err := r.db.Get(&p, "SELECT * FROM products WHERE id = $1", id)
	return &p, err
}

func (r *productRepository) GetAll() ([]*domain.Product, error) {
	var products []*domain.Product
	err := r.db.Select(&products, "SELECT * FROM products ORDER BY created_at DESC")
	return products, err
}

func (r *productRepository) Update(p *domain.Product) error {
	query := `
		UPDATE products
		SET name=$1, description=$2, price=$3, category=$4, stock=$5, updated_at=NOW()
		WHERE id=$6`
	_, err := r.db.Exec(query, p.Name, p.Description, p.Price, p.Category, p.Stock, p.ID)
	return err
}

func (r *productRepository) Delete(id int64) error {
	_, err := r.db.Exec("DELETE FROM products WHERE id = $1", id)
	return err
}
