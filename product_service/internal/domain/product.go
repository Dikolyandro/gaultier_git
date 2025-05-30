package domain

type Product struct {
	ID          int64   `db:"id"`
	Name        string  `db:"name"`
	Description string  `db:"description"`
	Price       float64 `db:"price"`
	Category    string  `db:"category"`
	Stock       int     `db:"stock"`
	CreatedAt   string  `db:"created_at"`
	UpdatedAt   string  `db:"updated_at"`
}

type ProductRepository interface {
	Create(product *Product) error
	GetByID(id int64) (*Product, error)
	GetAll() ([]*Product, error)
	Update(product *Product) error
	Delete(id int64) error
}

type ProductUsecase interface {
	Create(product *Product) error
	GetByID(id int64) (*Product, error)
	GetAll() ([]*Product, error)
	Update(product *Product) error
	Delete(id int64) error
}
