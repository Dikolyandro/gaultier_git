package usecase

import "github.com/da-er-gaultier/delivery_service/internal/domain"

type deliveryUsecase struct {
	repo domain.DeliveryRepository
}

func NewDeliveryUsecase(r domain.DeliveryRepository) domain.DeliveryUsecase {
	return &deliveryUsecase{repo: r}
}

func (u *deliveryUsecase) Create(d *domain.Delivery) error {
	return u.repo.Create(d)
}

func (u *deliveryUsecase) UpdateStatus(id int, status string) error {
	return u.repo.UpdateStatus(id, status)
}

func (u *deliveryUsecase) GetByID(id int) (*domain.Delivery, error) {
	return u.repo.GetByID(id)
}

func (u *deliveryUsecase) GetByOrderID(orderID int) (*domain.Delivery, error) {
	return u.repo.GetByOrderID(orderID)
}
