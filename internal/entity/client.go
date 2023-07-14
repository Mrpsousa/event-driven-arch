package entity

import (
	"errors"
	"time"

	"github.com/google/uuid"
)

type Client struct {
	ID        string     `json:"id"`
	Name      string     `json:"name"`
	Email     string     `json:"email"`
	Accounts  []*Account `json:"accounts"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
}

func NewClient(name, email string) (*Client, error) {
	client := &Client{
		ID:        uuid.New().String(),
		Name:      name,
		Email:     email,
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
	}
	err := client.Validate()
	if err != nil {
		return nil, err
	}
	return client, nil

}

func (c *Client) Validate() error {
	if c.Name == "" {
		return errors.New("name_is_required")
	}
	if c.Email == "" {
		return errors.New("email_is_required")
	}
	return nil
}

func (c *Client) Update(name, email string) error {
	c.Name = name
	c.Email = email
	c.UpdatedAt = time.Now().UTC()
	err := c.Validate()
	if err != nil {
		return err
	}
	return nil
}

func (c *Client) AddAccount(account *Account) error {
	if account.Client.ID != c.ID {
		return errors.New("account_does_not_belong_to_client")
	}
	c.Accounts = append(c.Accounts, account)
	return nil
}
