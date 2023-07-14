package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateAccountSucess(t *testing.T) {
	client, _ := NewClient("jhon", "eu@email.com")
	account := NewAccount(client)
	assert.NotNil(t, account)
	assert.Equal(t, client.ID, account.Client.ID)
}

func TestCreateAccountEmptyClient(t *testing.T) {
	account := NewAccount(nil)
	assert.Nil(t, account)
}

func TestCreditAccountSucess(t *testing.T) {
	client, _ := NewClient("jhon", "eu@email.com")
	account := NewAccount(client)
	account.Credit(100)
	assert.Equal(t, float64(100), account.Balance)
}

func TestDebitAccountSucess(t *testing.T) {
	client, _ := NewClient("jhon", "eu@email.com")
	account := NewAccount(client)
	account.Credit(100)
	account.Debit(50)
	assert.Equal(t, float64(50), account.Balance)
}
