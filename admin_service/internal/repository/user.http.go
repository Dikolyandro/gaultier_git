package repository

import (
	"encoding/json"
	"fmt"
	"github.com/da-er-gaultier/admin_service/internal/domain"
	"net/http"
	"time"
)

type userHTTPClient struct {
	baseURL string
	client  *http.Client
}

func NewUserHTTPClient(baseURL string) domain.UserClient {
	return &userHTTPClient{
		baseURL: baseURL,
		client:  &http.Client{Timeout: 5 * time.Second},
	}
}

func (c *userHTTPClient) ListUsers() ([]domain.User, error) {
	resp, err := c.client.Get(fmt.Sprintf("%s/internal/users", c.baseURL))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var users []domain.User
	if err := json.NewDecoder(resp.Body).Decode(&users); err != nil {
		return nil, err
	}
	return users, nil
}

func (c *userHTTPClient) DeleteUser(id int) error {
	req, err := http.NewRequest(http.MethodDelete, fmt.Sprintf("%s/internal/users/%d", c.baseURL, id), nil)
	if err != nil {
		return err
	}
	resp, err := c.client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusNoContent {
		return fmt.Errorf("failed to delete user: status %s", resp.Status)
	}
	return nil
}
