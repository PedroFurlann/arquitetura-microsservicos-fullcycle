package gateway

import "github.com.br/PedroFurlann/arquitetura-microsservicos-fullcycle/walletcore/internal/entity"

type TransactionGateway interface {
	Create(transaction *entity.Transaction) error
}
