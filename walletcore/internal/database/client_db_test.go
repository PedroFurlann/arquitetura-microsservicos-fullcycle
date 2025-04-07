package database

import (
	"database/sql"
	"testing"

	"github.com.br/PedroFurlann/arquitetura-microsservicos-fullcycle/walletcore/internal/entity"
	_ "github.com/mattn/go-sqlite3"
	"github.com/stretchr/testify/suite"
)

type ClientDBTestSuite struct {
	suite.Suite
	db       *sql.DB
	clientDB *ClientDB
}

func (suite *ClientDBTestSuite) SetupTest() {
	db, err := sql.Open("sqlite3", ":memory:")
	suite.Nil(err)

	suite.db = db

	db.Exec("CREATE TABLE clients (id varchar(255), name varchar(255), email varchar(255), created_at date)")

	suite.clientDB = NewClientDB(db)
}

func (suite *ClientDBTestSuite) TearDownTest() {
	defer suite.db.Close()
	suite.db.Exec("DROP TABLE clients")
}

func TestClientDBTestSuite(t *testing.T) {
	suite.Run(t, new(ClientDBTestSuite))
}

func (suite *ClientDBTestSuite) TestSaveClient() {
	client := &entity.Client{
		ID:    "123",
		Name:  "Pedro",
		Email: "email@example.com",
	}

	err := suite.clientDB.Save(client)
	suite.Nil(err)
}

func (suite *ClientDBTestSuite) TestGetClient() {
	client, _ := entity.NewClient("Pedro", "email@example.com")
	suite.clientDB.Save(client)
	result, err := suite.clientDB.Get(client.ID)
	suite.Nil(err)
	suite.Equal(client.ID, result.ID)
	suite.Equal(client.Name, result.Name)
	suite.Equal(client.Email, result.Email)
}
