package gateway

import "github.com/matheuslch/fc-ms-wallet-core/internal/entity"

type TransactionGateway interface {
	Create(transaction *entity.Transaction) error
}
