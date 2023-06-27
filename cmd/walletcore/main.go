package main

import (
	"database/sql"
	"fmt"

	"github.com/clebersimm/fc-ms-wallet/internal/database"
	"github.com/clebersimm/fc-ms-wallet/internal/event"
	"github.com/clebersimm/fc-ms-wallet/internal/usecase/create_account"
	"github.com/clebersimm/fc-ms-wallet/internal/usecase/create_client"
	"github.com/clebersimm/fc-ms-wallet/internal/usecase/create_transaction"
	"github.com/clebersimm/fc-ms-wallet/pkg/events"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", "root", "root", "mysql", "3306", "wallet"))
	if err != nil {
		panic(err)
	}
	defer db.Close()
	eventDispatcher := events.NewEventDispatcher()
	transactionCreatedEvent := event.NewTransactionCreated()
	eventDispatcher.Register("TransactionCreated", handler)
	clientDb := database.NewClientDB(db)
	accountDb := database.NewAccountDB(db)
	transactionDb := database.NewTransactionDB(db)
	createClientUseCase := create_client.NewCreateClientUseCase(clientDb)
	createAccountUseCase := create_account.NewCreateAccountUse(accountDb, clientDb)
	createTransactionUseCase := create_transaction.NewCreateTransactionUseCase(transactionDb, accountDb, eventDispatcher, transactionCreatedEvent)
}
