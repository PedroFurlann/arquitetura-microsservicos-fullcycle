package database

import (
	"database/sql"
	"testing"

	"github.com.br/PedroFurlann/arquitetura-microsservicos-fullcycle/internal/entity"
	"github.com/stretchr/testify/suite"
)

type AccountDBTestSuite struct {
	suite.Suite
	db        *sql.DB
	accountDB *AccountDB
	client    *entity.Client
}

func (suite *AccountDBTestSuite) SetupTest() {
	db, err := sql.Open("sqlite3", ":memory:")
	suite.Nil(err)
	suite.db = db
	db.Exec("CREATE TABLE clients (id varchar(255), name varchar(255), email varchar(255), created_at date)")
	db.Exec("CREATE TABLE accounts (id varchar(255), client_id varchar(255), balance float, created_at date)")
	suite.accountDB = NewAccountDB(db)
	suite.client, _ = entity.NewClient("Pedro", "email@example.com")
}

func (suite *AccountDBTestSuite) TearDownTest() {
	defer suite.db.Close()
	suite.db.Exec("DROP TABLE clients")
	suite.db.Exec("DROP TABLE accounts")
}

func TestAccountDBTestSuite(t *testing.T) {
	suite.Run(t, new(AccountDBTestSuite))
}

func (suite *AccountDBTestSuite) TestSaveAccount() {
	account := entity.NewAccount(suite.client)
	err := suite.accountDB.Save(account)
	suite.Nil(err)
}

func (suite *AccountDBTestSuite) TestAccountFindByID() {
	suite.db.Exec("INSERT INTO clients (id, name, email, created_at) VALUES (?, ?, ?, ?)", suite.client.ID, suite.client.Name, suite.client.Email, suite.client.CreatedAt)
	account := entity.NewAccount(suite.client)
	err := suite.accountDB.Save(account)
	suite.Nil(err)

	result, err := suite.accountDB.FindByID(account.ID)
	suite.Nil(err)
	suite.Equal(account.ID, result.ID)
	suite.Equal(account.Balance, result.Balance)
	suite.Equal(account.Client.ID, result.Client.ID)
	suite.Equal(account.Client.Name, result.Client.Name)
	suite.Equal(account.Client.Email, result.Client.Email)
}
