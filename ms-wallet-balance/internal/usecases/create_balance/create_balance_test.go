package create_balance

import (
	"testing"
	"time"

	"github.com/matheuslch/fc-ms-wallet-balance/internal/entity"
	"github.com/matheuslch/fc-ms-wallet-balance/internal/usecases/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestCreateBalanceUseCase_Execute(t *testing.T) {
	balanceRepository := mocks.NewBalanceGatewayMock()
	uc := NewCreateBalanceUseCase(balanceRepository)

	input := CreateBalanceInputDTO{
		AccountIDFrom:        "1",
		AccountIDTo:          "2",
		BalanceAccountIDFrom: 10,
		BalanceAccountIDTo:   20,
		UpdateAt:             "2021-01-01T00:00:00Z",
	}
	balanceRepository.On("FindByAccountID", mock.Anything).Return(&entity.Balance{}, nil).Times(2)
	balanceRepository.On("Save", mock.Anything).Return(nil)

	err := uc.Execute(input)
	assert.Nil(t, err)
}

func TestCreateBalanceUseCase_ExecuteWithOutdatedFrom(t *testing.T) {
	balanceRepository := mocks.NewBalanceGatewayMock()
	uc := NewCreateBalanceUseCase(balanceRepository)

	input := CreateBalanceInputDTO{
		AccountIDFrom:        "1",
		AccountIDTo:          "2",
		BalanceAccountIDFrom: 10,
		BalanceAccountIDTo:   20,
		UpdateAt:             "2023-01-01T00:00:00Z",
	}

	balanceFrom := entity.Balance{
		AccountID:     "1",
		BalanceAmount: 5,
		UpdatedAt:     time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC),
	}
	balanceRepository.On("FindByAccountID", "1").Return(&balanceFrom, nil)
	balanceRepository.On("FindByAccountID", "2").Return(&balanceFrom, nil)

	err := uc.Execute(input)
	assert.EqualError(t, err, "the balance from is outdated")
}

func TestCreateBalanceUseCase_ExecuteWithOutdatedTo(t *testing.T) {
	balanceTo := entity.Balance{
		AccountID:     "2",
		BalanceAmount: 5,
		UpdatedAt:     time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC),
	}

	balanceRepository := mocks.NewBalanceGatewayMock()
	balanceRepository.On("FindByAccountID", mock.Anything).Return(&entity.Balance{}, nil).Times(1)
	balanceRepository.On("FindByAccountID", "2").Return(&balanceTo, nil).Times(1)
	balanceRepository.On("Save", mock.Anything).Return(nil)

	uc := NewCreateBalanceUseCase(balanceRepository)
	input := CreateBalanceInputDTO{
		AccountIDFrom:        "1",
		AccountIDTo:          "2",
		BalanceAccountIDFrom: 10,
		BalanceAccountIDTo:   20,
		UpdateAt:             "2023-01-01T00:00:00Z",
	}

	err := uc.Execute(input)
	assert.EqualError(t, err, "the balance to is outdated")
}
