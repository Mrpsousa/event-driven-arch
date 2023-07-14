package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateTransactionSucess(t *testing.T) {
	client, _ := NewClient("jhon", "eu@email.com")
	account := NewAccount(client)
	client2, _ := NewClient("jhon2", "eu2@email.com")
	account2 := NewAccount(client2)

	account.Credit(1000)
	account2.Credit(1000)

	transaction, err := NewTransaction(account, account2, 100)
	assert.Nil(t, err)
	assert.NotNil(t, transaction)
	assert.Equal(t, account.ID, transaction.AccountFrom.ID)
	assert.Equal(t, 1100.0, account2.Balance)
	assert.Equal(t, 900.0, account.Balance)

}

func TestCreateTransactionWithInsuficientBalance(t *testing.T) {
	client, _ := NewClient("jhon", "eu@email.com")
	account := NewAccount(client)
	client2, _ := NewClient("jhon2", "eu2@email.com")
	account2 := NewAccount(client2)

	account.Credit(1000)
	account2.Credit(1000)

	transaction, err := NewTransaction(account, account2, 2000)
	assert.NotNil(t, err)
	assert.Nil(t, transaction)
	assert.Equal(t, err.Error(), "insufficient_funds")
	assert.Equal(t, 1000.0, account2.Balance)
	assert.Equal(t, 1000.0, account.Balance)

}
