package find_balance

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
	uc := NewFindBalanceUseCase(balanceRepository)

	balance := entity.Balance{
		AccountID:     "1234",
		BalanceAmount: 5,
		UpdatedAt:     time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC),
	}
	balanceRepository.On("FindByAccountID", mock.Anything).Return(&balance, nil)

	findBalance, err := uc.Execute("1234")
	assert.Nil(t, err)
	assert.Equal(t, "1234", findBalance.AccountID)
}

func TestCreateBalanceUseCase_ExecuteWithNotFound(t *testing.T) {
	balanceRepository := mocks.NewBalanceGatewayMock()
	uc := NewFindBalanceUseCase(balanceRepository)

	balanceRepository.On("FindByAccountID", mock.Anything).Return(&entity.Balance{}, nil)

	findBalance, err := uc.Execute("1234")
	assert.Nil(t, findBalance)
	assert.EqualError(t, err, "AccountID not found")
}
