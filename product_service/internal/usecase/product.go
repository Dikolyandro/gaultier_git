package usecase

import (
	"errors"
	"product_service/internal/domain"
)

type productUsecase struct {
	repo domain.ProductRepository
}

func NewProductUsecase(r domain.ProductRepository) domain.ProductUsecase {
	return &productUsecase{repo: r}
}

func (u *productUsecase) Create(p *domain.Product) error {
	if p.Name == "" || p.Price <= 0 {
		return errors.New("invalid product data")
	}
	return u.repo.Create(p)
}

func (u *productUsecase) GetByID(id int64) (*domain.Product, error) {
	return u.repo.GetByID(id)
}

func (u *productUsecase) GetAll() ([]*domain.Product, error) {
	return u.repo.GetAll()
}

func (u *productUsecase) Update(p *domain.Product) error {
	if p.ID == 0 {
		return errors.New("missing product ID")
	}
	return u.repo.Update(p)
}

func (u *productUsecase) Delete(id int64) error {
	if id == 0 {
		return errors.New("invalid product ID")
	}
	return u.repo.Delete(id)
}
