package createaccount

import (
	"github.com.br/PedroFurlann/arquitetura-microsservicos-fullcycle/internal/entity"
	"github.com.br/PedroFurlann/arquitetura-microsservicos-fullcycle/internal/gateway"
)

type CreateAccountInputDTO struct {
	ClientID string
}

type CreateAccountOutputDTO struct {
	ID string
}

type CreateAccountUseCase struct {
	AccountGateway gateway.AccountGateway
	ClientGateway  gateway.ClientGateway
}

func NewCreateAccountUseCase(accountGateway gateway.AccountGateway, clientGateway gateway.ClientGateway) *CreateAccountUseCase {
	return &CreateAccountUseCase{
		AccountGateway: accountGateway,
		ClientGateway:  clientGateway,
	}
}

func (uc *CreateAccountUseCase) Execute(input CreateAccountInputDTO) (*CreateAccountOutputDTO, error) {
	client, err := uc.ClientGateway.Get(input.ClientID)
	if err != nil {
		return nil, err
	}

	account := entity.NewAccount(client)

	err = uc.AccountGateway.Save(account)
	if err != nil {
		return nil, err
	}

	output := &CreateAccountOutputDTO{
		ID: account.ID,
	}

	return output, nil
}
