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

func TestUpdateClient(t *testing.T) {
	client, _ := NewClient("Pedro Furlan", "email@example.com")
	err := client.Update("Pedro Furlan Updated", "email@exampleupdated.com")
	assert.Nil(t, err)
	assert.Equal(t, "Pedro Furlan Updated", client.Name)
	assert.Equal(t, "email@exampleupdated.com", client.Email)
}

func TestUpdateClientWhenArgsAreInvalid(t *testing.T) {
	client, _ := NewClient("Pedro Furlan", "email@example.com")
	err := client.Update("", "")
	assert.Error(t, err, "name is required")
	assert.Error(t, err, "email is required")
}
