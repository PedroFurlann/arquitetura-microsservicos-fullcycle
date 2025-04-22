package gateway

import "github.com.br/PedroFurlann/arquitetura-microsservicos-fullcycle/internal/entity"

type TransactionGateway interface {
	Create(transaction *entity.Transaction) error
}
