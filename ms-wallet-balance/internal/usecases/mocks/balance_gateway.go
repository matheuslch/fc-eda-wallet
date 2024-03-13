package mocks

import (
	"github.com/matheuslch/fc-ms-wallet-balance/internal/entity"
	"github.com/stretchr/testify/mock"
)

type BalanceGatewayMock struct {
	mock.Mock
}

func NewBalanceGatewayMock() *BalanceGatewayMock {
	return &BalanceGatewayMock{}
}

func (m *BalanceGatewayMock) Save(balance *entity.Balance) error {
	args := m.Called(balance)
	return args.Error(0)
}

func (m *BalanceGatewayMock) FindByAccountID(AccountID string) (*entity.Balance, error) {
	args := m.Called(AccountID)
	return args.Get(0).(*entity.Balance), args.Error(1)
}

func (m *BalanceGatewayMock) Update(balance *entity.Balance) error {
	args := m.Called(balance)
	return args.Error(0)
}
