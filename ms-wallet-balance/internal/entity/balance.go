package entity

import (
	"errors"
	"time"

	"github.com/google/uuid"
)

type Balance struct {
	ID            string
	AccountID     string
	BalanceAmount float64
	UpdatedAt     time.Time
}

func NewBalance(accountID string, balanceAmount float64, updatedAT time.Time) (*Balance, error) {
	balance := &Balance{
		ID:            uuid.New().String(),
		AccountID:     accountID,
		BalanceAmount: balanceAmount,
		UpdatedAt:     updatedAT,
	}

	err := balance.Validate()
	if err != nil {
		return nil, err
	}

	return balance, nil
}

func (t *Balance) Validate() error {
	if t.AccountID == "" {
		return errors.New("AccountID cannot be empty")
	}

	if t.UpdatedAt.IsZero() {
		return errors.New("UpdatedAt cannot be empty")
	}

	return nil
}
