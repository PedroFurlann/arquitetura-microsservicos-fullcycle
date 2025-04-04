package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateNewClient(t *testing.T) {
	client, err := NewClient("Pedro Furlan", "email@example.com")
	assert.Nil(t, err)
	assert.NotNil(t, client)
	assert.Equal(t, "Pedro Furlan", client.Name)
	assert.Equal(t, "email@example.com", client.Email)
}

func TestCreateNewClientWhenArgsAreInvalid(t *testing.T) {
	client, err := NewClient("", "")
	assert.NotNil(t, err)
	assert.Nil(t, client)
}
