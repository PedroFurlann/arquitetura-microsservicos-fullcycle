package gateway

import "github.com.br/PedroFurlann/arquitetura-microsservicos-fullcycle/walletcore/internal/entity"

type AccountGateway interface {
	Save(account *entity.Account) error
	FindByID(id string) (*entity.Account, error)
}
