package createclient

import (
	"testing"

	"github.com.br/PedroFurlann/arquitetura-microsservicos-fullcycle/internal/entity"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockClientGateway struct {
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

func TestCreateClientUseCase_Execute(t *testing.T) {
	m := &MockClientGateway{}
	m.On("Save", mock.Anything).Return(nil)

	useCase := NewCreateClientUseCase(m)
	output, err := useCase.Execute(CreateClientInputDTO{
		Name:  "Pedro Furlan",
		Email: "email@example.com",
	})
	assert.Nil(t, err)
	assert.NotNil(t, output)
	assert.NotEmpty(t, output.ID)
	assert.Equal(t, "Pedro Furlan", output.Name)
	assert.Equal(t, "email@example.com", output.Email)

	m.AssertExpectations(t)
	m.AssertNumberOfCalls(t, "Save", 1)
}
