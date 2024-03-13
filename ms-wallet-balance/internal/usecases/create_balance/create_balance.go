package create_balance

import (
	"errors"
	"time"

	"github.com/matheuslch/fc-ms-wallet-balance/internal/entity"
	"github.com/matheuslch/fc-ms-wallet-balance/internal/gateway"
)

type ConsumerDTO struct {
	Name    string                `json:"Name"`
	Payload CreateBalanceInputDTO `json:"Payload"`
}

type CreateBalanceInputDTO struct {
	AccountIDFrom        string  `json:"account_id_from"`
	AccountIDTo          string  `json:"account_id_to"`
	BalanceAccountIDFrom float64 `json:"balance_account_id_from"`
	BalanceAccountIDTo   float64 `json:"balance_account_id_to"`
	UpdateAt             string  `json:"updated_at"`
}

type CreateBalanceUseCase struct {
	BalanceGateway gateway.BalanceGateway
}

func NewCreateBalanceUseCase(balanceGateway gateway.BalanceGateway) *CreateBalanceUseCase {
	return &CreateBalanceUseCase{BalanceGateway: balanceGateway}
}

func (uc *CreateBalanceUseCase) Execute(input CreateBalanceInputDTO) error {

	updatedAt, err := time.Parse(time.RFC3339, input.UpdateAt)
	if err != nil {
		return err
	}

	findBalanceFrom, err := uc.BalanceGateway.FindByAccountID(input.AccountIDFrom)
	if err != nil {
		return err
	}
	findBalanceTo, err := uc.BalanceGateway.FindByAccountID(input.AccountIDTo)
	if err != nil {
		return err
	}

	if findBalanceFrom.AccountID == input.AccountIDFrom {
		if findBalanceFrom.UpdatedAt.After(updatedAt) {
			return errors.New("the balance from is outdated")
		}

		findBalanceFrom.BalanceAmount = input.BalanceAccountIDFrom
		findBalanceFrom.UpdatedAt = updatedAt
		err = uc.BalanceGateway.Update(findBalanceFrom)
		if err != nil {
			return err
		}
	} else {
		balanceFrom, err := entity.NewBalance(input.AccountIDFrom, input.BalanceAccountIDFrom, updatedAt)
		if err != nil {
			return err
		}
		err = uc.BalanceGateway.Save(balanceFrom)
		if err != nil {
			return err
		}
	}

	if findBalanceTo.AccountID == input.AccountIDTo {
		if findBalanceTo.UpdatedAt.After(updatedAt) {
			return errors.New("the balance to is outdated")
		}

		findBalanceTo.BalanceAmount = input.BalanceAccountIDTo
		findBalanceTo.UpdatedAt = updatedAt
		err = uc.BalanceGateway.Update(findBalanceTo)
		if err != nil {
			return err
		}
	} else {
		balanceTo, err := entity.NewBalance(input.AccountIDTo, input.BalanceAccountIDTo, updatedAt)
		if err != nil {
			return err
		}
		err = uc.BalanceGateway.Save(balanceTo)
		if err != nil {
			return err
		}
	}
	return nil
}
