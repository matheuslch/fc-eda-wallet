package entity

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateNewClient(t *testing.T) {
	client, err := NewClient("John Doe", "j@doe.com")
	assert.Nil(t, err)
	assert.NotNil(t, client)
	assert.Equal(t, "John Doe", client.Name)
	assert.Equal(t, "j@doe.com", client.Email)
}

func TestCreateNewClientWhenArgsAreInvalid(t *testing.T) {
	client, err := NewClient("", "")
	assert.NotNil(t, err)
	assert.Nil(t, client)
}

func TestUpdateClient(t *testing.T) {
	client, _ := NewClient("John Doe", "j@doe.com")
	err := client.Update("John Updated", "john@doe.com")
	assert.Nil(t, err)
	assert.Equal(t, "John Updated", client.Name)
	assert.Equal(t, "john@doe.com", client.Email)
}

func TestUpdateClientWithInvalidArgs(t *testing.T) {
	client, _ := NewClient("John Doe", "j@doe.com")
	err := client.Update("", "john@doe.com")
	assert.Error(t, err, "name is required")
}

func TestAddAccountToClient(t *testing.T) {
	client, _ := NewClient("John Doe", "john@doe.com")
	account := NewAccount(client)
	err := client.AddAccount(account)
	assert.Nil(t, err)
	assert.Equal(t, 1, len(client.Accounts))
}

func TestAddAccountWithError(t *testing.T) {
	client, _ := NewClient("John Doe", "john@doe.com")
	account := NewAccount(client)
	client.AddAccount(account)

	client2, _ := NewClient("John Doe 2", "john2@doe.com")
	account2 := NewAccount(client2)

	err := client.AddAccount(account2)
	assert.NotNil(t, err)
	assert.Equal(t, errors.New("accounts does not belong to this client"), err)
}
