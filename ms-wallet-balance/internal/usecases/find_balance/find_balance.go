package find_balance

import (
	"errors"
	"time"

	"github.com/matheuslch/fc-ms-wallet-balance/internal/gateway"
)

type FindBalanceInputDTO struct {
	AccountID string  `json:"account_id"`
	Balance   float64 `json:"balance"`
	UpdateAt  string  `json:"updated_at"`
}

type FindBalanceUseCase struct {
	BalanceGateway gateway.BalanceGateway
}

func NewFindBalanceUseCase(balanceGateway gateway.BalanceGateway) *FindBalanceUseCase {
	return &FindBalanceUseCase{BalanceGateway: balanceGateway}
}

func (uc *FindBalanceUseCase) Execute(AccountID string) (*FindBalanceInputDTO, error) {

	findBalanceFrom, err := uc.BalanceGateway.FindByAccountID(AccountID)
	if err != nil {
		return nil, err
	}

	if findBalanceFrom == nil || findBalanceFrom.AccountID == "" {
		return nil, errors.New("AccountID not found")
	}

	return &FindBalanceInputDTO{
		AccountID: findBalanceFrom.AccountID,
		Balance:   findBalanceFrom.BalanceAmount,
		UpdateAt:  findBalanceFrom.UpdatedAt.Format(time.RFC3339),
	}, nil
}
