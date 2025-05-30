package usecase

import (
	"github.com/da-er-gaultier/cart_service/internal/domain"
)

type cartUsecase struct {
	repo domain.CartRepository
}

func NewCartUsecase(r domain.CartRepository) domain.CartUsecase {
	return &cartUsecase{repo: r}
}

func (u *cartUsecase) AddToCart(item *domain.CartItem) error {
	return u.repo.Add(item)
}

func (u *cartUsecase) GetCart(userID int) ([]domain.CartItem, error) {
	return u.repo.GetByUser(userID)
}

func (u *cartUsecase) RemoveFromCart(itemID int) error {
	return u.repo.Delete(itemID)
}
