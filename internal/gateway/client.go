package gateway

import "github.com.br/PedroFurlann/arquitetura-microsservicos-fullcycle/internal/entity"

type ClientGateway interface {
	Get(id string) (*entity.Client, error)
	Save(client *entity.Client) error
}
