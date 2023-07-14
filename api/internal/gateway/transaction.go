package gateway

import "github.com/mrpsousa/internal/entity"

type TransactionGateway interface {
	Create(transaction *entity.Transaction) error
}
