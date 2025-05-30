package usecase

import (
	"github.com/da-er-gaultier/admin_service/internal/domain"
)

type adminUsecase struct {
	repo       domain.AdminRepository
	userClient domain.UserClient
}

func New(repo domain.AdminRepository, userClient domain.UserClient) domain.AdminUsecase {
	return &adminUsecase{
		repo:       repo,
		userClient: userClient,
	}
}

// 👇 теперь работает через userService
func (u *adminUsecase) ListUsers() ([]domain.User, error) {
	return u.userClient.ListUsers()
}

func (u *adminUsecase) DeleteUser(userID int) error {
	return u.userClient.DeleteUser(userID)
}

// всё остальное — через локальный репозиторий
func (u *adminUsecase) ListProducts() ([]domain.Product, error) {
	return u.repo.GetAllProducts()
}

func (u *adminUsecase) DeleteProduct(productID int) error {
	return u.repo.DeleteProductByID(productID)
}

func (u *adminUsecase) ListOrders() ([]domain.Order, error) {
	return u.repo.GetAllOrders()
}

func (u *adminUsecase) UpdateOrderStatus(orderID int, status string) error {
	return u.repo.UpdateOrderStatus(orderID, status)
}
