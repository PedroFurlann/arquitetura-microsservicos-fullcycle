package createtransaction

import (
	"github.com.br/PedroFurlann/arquitetura-microsservicos-fullcycle/internal/entity"
	"github.com.br/PedroFurlann/arquitetura-microsservicos-fullcycle/internal/gateway"
	"github.com.br/PedroFurlann/arquitetura-microsservicos-fullcycle/pkg/events"
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
	EventDispatcher    events.EventDispatcherInterface
	TransactionCreated events.EventInterface
}

func NewCreateTransactionUseCase(transactionGateway gateway.TransactionGateway, accountGateway gateway.AccountGateway, eventDispatcher events.EventDispatcherInterface, transactionCreated events.EventInterface) *CreateTransactionUseCase {
	return &CreateTransactionUseCase{
		TransactionGateway: transactionGateway,
		AccountGateway:     accountGateway,
		EventDispatcher:    eventDispatcher,
		TransactionCreated: transactionCreated,
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

	uc.TransactionCreated.SetPayload(output)
	uc.EventDispatcher.Dispatch(uc.TransactionCreated)

	return output, nil
}
