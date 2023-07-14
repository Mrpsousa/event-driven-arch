package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateNewClientSucess(t *testing.T) {
	client, err := NewClient("jhon", "eu@email.com")
	assert.NotNil(t, client)
	assert.NoError(t, err)
	assert.Equal(t, "jhon", client.Name)
	assert.Equal(t, "eu@email.com", client.Email)
}

func TestCreateNewClientWithOutName(t *testing.T) {
	client, err := NewClient("", "eu@email.com")
	assert.Nil(t, client)
	assert.NotNil(t, err)
	assert.Equal(t, "name_is_required", err.Error())

}

func TestCreateNewClientWithOutEmail(t *testing.T) {
	client, err := NewClient("jhon", "")
	assert.Nil(t, client)
	assert.NotNil(t, err)
	assert.Equal(t, "email_is_required", err.Error())
}

func TestUpdateClient(t *testing.T) {
	client, _ := NewClient("jhon", "eu@email.com")
	err := client.Update("jhon Update", "eu@email.com")
	assert.Nil(t, err)
	assert.Equal(t, "jhon Update", client.Name)
	assert.Equal(t, "eu@email.com", client.Email)
}

func TestUpdateClientWhenError(t *testing.T) {
	client, _ := NewClient("jhon", "eu@email.com")
	err := client.Update("", "eu@email.com")
	assert.NotNil(t, err)
	assert.Equal(t, "name_is_required", err.Error())
}

func TestAddAccountToClientSucess(t *testing.T) {
	client, err := NewClient("jhon", "eu@email.com")
	account := NewAccount(client)
	err = client.AddAccount(account)
	assert.Nil(t, err)
	assert.Equal(t, 1, len(client.Accounts))
}
