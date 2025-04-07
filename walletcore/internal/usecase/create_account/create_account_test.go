package createaccount

import (
	"testing"

	"github.com.br/PedroFurlann/arquitetura-microsservicos-fullcycle/walletcore/internal/entity"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockClientGateway struct {
	mock.Mock
}

type MockAccountGateway struct {
	mock.Mock
}

func (m *MockClientGateway) Get(id string) (*entity.Client, error) {
	args := m.Called(id)
	return args.Get(0).(*entity.Client), args.Error(1)
}

func (m *MockClientGateway) Save(client *entity.Client) error {
	args := m.Called(client)
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

func TestCreateAccountUseCase_Execute(t *testing.T) {
	client, _ := entity.NewClient("Pedro Furlan", "email@example.com")
	clientMock := &MockClientGateway{}
	clientMock.On("Get", client.ID).Return(client, nil)

	accountMock := &MockAccountGateway{}
	accountMock.On("Save", mock.Anything).Return(nil)

	uc := NewCreateAccountUseCase(accountMock, clientMock)
	output, err := uc.Execute(CreateAccountInputDTO{
		ClientID: client.ID,
	})
	assert.Nil(t, err)
	assert.NotNil(t, output.ID)
	clientMock.AssertExpectations(t)
	clientMock.AssertNumberOfCalls(t, "Get", 1)
	accountMock.AssertExpectations(t)
	accountMock.AssertNumberOfCalls(t, "Save", 1)
}
