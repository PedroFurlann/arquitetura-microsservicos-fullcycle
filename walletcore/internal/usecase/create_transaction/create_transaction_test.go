package createtransaction

import (
	"testing"

	"github.com.br/PedroFurlann/arquitetura-microsservicos-fullcycle/walletcore/internal/entity"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockTransactionGateway struct {
	mock.Mock
}

type MockAccountGateway struct {
	mock.Mock
}

func (m *MockTransactionGateway) Create(transaction *entity.Transaction) error {
	args := m.Called(transaction)
	return args.Error(0)
}

func (m *MockAccountGateway) Save(account *entity.Account) error {
	args := m.Called(account)
	return args.Error(0)
}

func (m *MockAccountGateway) FindByID(id string) (*entity.Account, error) {
	args := m.Called(id)
	return args.Get(0).(*entity.Account), args.Error(1)
}

func TestCreateTransactionUseCase_Execute(t *testing.T) {
	client1, _ := entity.NewClient("Client1", "Client1@example.com")
	client2, _ := entity.NewClient("Client2", "Client2@example.com")

	account1 := entity.NewAccount(client1)
	account2 := entity.NewAccount(client2)

	account1.Credit(1000)
	account2.Credit(1000)

	accountMock := &MockAccountGateway{}
	accountMock.On("FindByID", account1.ID).Return(account1, nil)
	accountMock.On("FindByID", account2.ID).Return(account2, nil)

	transactionMock := &MockTransactionGateway{}
	transactionMock.On("Create", mock.Anything).Return(nil)

	uc := NewCreateTransactionUseCase(transactionMock, accountMock)

	output, err := uc.Execute(CreateTransactionInputDTO{
		AccountFromID: account1.ID,
		AccountToID:   account2.ID,
		Amount:        100,
	})

	assert.Nil(t, err)
	assert.NotNil(t, output.ID)
	accountMock.AssertExpectations(t)
	accountMock.AssertNumberOfCalls(t, "FindByID", 2)
	transactionMock.AssertExpectations(t)
	transactionMock.AssertNumberOfCalls(t, "Create", 1)
}
