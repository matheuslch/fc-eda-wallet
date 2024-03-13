package mocks

import (
	"github.com/matheuslch/fc-ms-wallet-core/internal/entity"
	"github.com/stretchr/testify/mock"
)

type TransactionGatewayMock struct {
	mock.Mock
}

func NewTransactionGatewayMock() *TransactionGatewayMock {
	return &TransactionGatewayMock{}
}

func (m *TransactionGatewayMock) Create(transaction *entity.Transaction) error {
	args := m.Called(transaction)
	return args.Error(0)
}
