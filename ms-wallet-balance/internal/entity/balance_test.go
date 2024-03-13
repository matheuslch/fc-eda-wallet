package entity

import (
	"errors"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestCreateBalance(t *testing.T) {
	balance, err := NewBalance("1", 20, time.Now())
	assert.Nil(t, err)
	assert.NotNil(t, balance)
	assert.Equal(t, "1", balance.AccountID)
	assert.Equal(t, float64(20), balance.BalanceAmount)
}

func TestCreateBalanceWithAccountIdIsEmpty(t *testing.T) {
	balance, err := NewBalance("", 10, time.Now())
	assert.NotNil(t, err)
	assert.Equal(t, err, errors.New("AccountID cannot be empty"))
	assert.Nil(t, balance)
}
