package domain

type UserClient interface {
	ListUsers() ([]User, error)
	DeleteUser(id int) error
}
