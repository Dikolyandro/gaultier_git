package usecase

import "github.com/da-er-gaultier/order_service/internal/domain"

type orderUsecase struct {
	repo domain.OrderRepository
}

func NewOrderUsecase(r domain.OrderRepository) domain.OrderUsecase {
	return &orderUsecase{repo: r}
}

func (u *orderUsecase) Create(order *domain.Order) error {
	return u.repo.Create(order)
}

func (u *orderUsecase) Get(orderID int) (*domain.Order, error) {
	return u.repo.GetByID(orderID)
}

func (u *orderUsecase) GetByUser(userID int) ([]domain.Order, error) {
	return u.repo.GetByUserID(userID)
}

func (u *orderUsecase) Delete(orderID int) error {
	return u.repo.Delete(orderID)
}
