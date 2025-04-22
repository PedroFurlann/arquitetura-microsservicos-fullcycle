package main

import (
	"database/sql"
	"fmt"

	"github.com.br/PedroFurlann/arquitetura-microsservicos-fullcycle/internal/database"
	"github.com.br/PedroFurlann/arquitetura-microsservicos-fullcycle/internal/event"
	createaccount "github.com.br/PedroFurlann/arquitetura-microsservicos-fullcycle/internal/usecase/create_account"
	createclient "github.com.br/PedroFurlann/arquitetura-microsservicos-fullcycle/internal/usecase/create_client"
	createtransaction "github.com.br/PedroFurlann/arquitetura-microsservicos-fullcycle/internal/usecase/create_transaction"
	"github.com.br/PedroFurlann/arquitetura-microsservicos-fullcycle/pkg/events"
)

func main() {
	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", "root", "root", "mysql", "3306", "wallet"))
	if err != nil {
		panic(err)
	}

	defer db.Close()

	eventDispatcher := events.NewEventDispatcher()
	transactionCreatedEvent := event.NewTransactionCreated()
	// eventDispatcher.Register("TransactionCreated")

	clientDb := database.NewClientDB(db)
	accountDb := database.NewAccountDB(db)
	transactionDb := database.NewTransactionDB(db)

	createClientUseCase := createclient.NewCreateClientUseCase(clientDb)
	createAccountUseCase := createaccount.NewCreateAccountUseCase(accountDb, clientDb)
	createTransactionUseCase := createtransaction.NewCreateTransactionUseCase(transactionDb, accountDb, eventDispatcher, transactionCreatedEvent)
}
