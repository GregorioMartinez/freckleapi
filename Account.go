package freckleapi

import (
	"encoding/json"
)

type Account struct {
	CreatedAt        string  `json:"created_at"`
	ID               float64 `json:"id"`
	InvoicingEnabled bool    `json:"invoicing_enabled"`
	Name             string  `json:"name"`
	Owner            User    `json:"owner"`
	UpdatedAt        string  `json:"updated_at"`
	URL              string  `json:"url"`
}

type AccountService struct {
	client *Client
}

func NewAccountService(client *Client) *AccountService {
	return &AccountService{client}
}

type AccountListCall struct {
	service *AccountService
}

func (s *AccountService) List() *AccountListCall {
	return &AccountListCall{
		service: s,
	}
}

func (c *AccountListCall) Do() (*Account, error) {
	data, err := c.service.client.run("GET", "account", nil)
	if err != nil {
		return nil, err
	}

	var account Account
	err = json.Unmarshal(data, &account)
	if err != nil {
		return nil, err
	}

	return &account, nil
}
