package consul_test

import (
	"testing"

	consul "github.com/segurosfalabella/imperium-backinator/backuper/backuper-consul"
	"github.com/stretchr/testify/assert"
)

func TestShouldReturnErrorWhenEndpointIsNotValid(t *testing.T) {
	var endpoint = "hola mundo"
	consul := &consul.Backuper{
		Endpoint: endpoint,
	}

	err := consul.Backup()

	assert.NotNil(t, err)
}

func TestShouldReturnErrorWhenTokenIsNotDefined(t *testing.T) {
	var token = ""
	var endpoint = "http://www.google.com"
	consul := &consul.Backuper{
		Token:    token,
		Endpoint: endpoint,
	}

	err := consul.Backup()

	assert.NotNil(t, err)
}

func TestShouldValidateMethodReturnTrueWhenEndpointIsValid(t *testing.T) {
	var endpoint = "http://www.google.com"

	state := consul.ValidateEndpoint(endpoint)

	assert.True(t, state)
}

func TestShoudlValidateMethodReturnFalseWhenEndpointIsInvalid(t *testing.T) {
	var endpoint = "hola mundo"

	state := consul.ValidateEndpoint(endpoint)

	assert.False(t, state)
}
