package createtransaction

import (
	"github.com.br/PedroFurlann/arquitetura-microsservicos-fullcycle/walletcore/internal/entity"
	"github.com.br/PedroFurlann/arquitetura-microsservicos-fullcycle/walletcore/internal/gateway"
)

type CreateTransactionInputDTO struct {
	AccountFromID string
	AccountToID   string
	Amount        float64
}

type CreateTransactionOutputDTO struct {
	ID string
}

type CreateTransactionUseCase struct {
	TransactionGateway gateway.TransactionGateway
	AccountGateway     gateway.AccountGateway
}

func NewCreateTransactionUseCase(transactionGateway gateway.TransactionGateway, accountGateway gateway.AccountGateway) *CreateTransactionUseCase {
	return &CreateTransactionUseCase{
		TransactionGateway: transactionGateway,
		AccountGateway:     accountGateway,
	}
}

func (uc *CreateTransactionUseCase) Execute(input CreateTransactionInputDTO) (*CreateTransactionOutputDTO, error) {
	accountFrom, err := uc.AccountGateway.FindByID(input.AccountFromID)
	if err != nil {
		return nil, err
	}

	accountTo, err := uc.AccountGateway.FindByID(input.AccountToID)
	if err != nil {
		return nil, err
	}

	transaction, err := entity.NewTransaction(accountFrom, accountTo, input.Amount)
	if err != nil {
		return nil, err
	}

	err = uc.TransactionGateway.Create(transaction)
	if err != nil {
		return nil, err
	}

	output := &CreateTransactionOutputDTO{
		ID: transaction.ID,
	}

	return output, nil
}
