package consul_test

import (
	"errors"
	"io"
	"testing"

	"github.com/hashicorp/consul/api"
	consul "github.com/segurosfalabella/imperium-backinator/backuper/backuper-consul"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type mockClient struct {
	mock.Mock
}

type Snapshot struct {
	c *mockClient
}

// Snapshot returns a handle that exposes the snapshot endpoints.
func (c *mockClient) Snapshot() *Snapshot {
	return &Snapshot{c}
}

func (s *Snapshot) Save(q *api.QueryOptions) (io.ReadCloser, *api.QueryMeta, error) {
	err := errors.New("could not save snapshot")
	return nil, nil, err
}

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

func TestShoudlReturnErrorWhenCallNewClientFail(t *testing.T) {
	var endpoint = "http://www.google.com"
	var token = "token-demo"
	consulInstance := &consul.Backuper{
		Endpoint: endpoint,
		Token:    token,
	}
	var oldApi = consul.API
	defer func() { consul.API = oldApi }()
	consul.API = func(config *api.Config) (*api.Client, error) {
		return nil, errors.New("could not create new client")
	}

	err := consulInstance.Backup()

	assert.NotNil(t, err)
	assert.Equal(t, err.Error(), "could not create new client")
}

// func TestShouldReturnErrorWhenSnapshotSaveFail(t *testing.T) {
// 	var endpoint = "http://www.google.com"
// 	var token = "token-demo"
// 	consulInstance := &consul.Backuper{
// 		Endpoint: endpoint,
// 		Token:    token,
// 	}
// 	var oldApi = consul.API
// 	defer func() { consul.API = oldApi }()
// 	consul.API = func(config *api.Config) (*api.Client, error) {
// 		client := new(api.Client)
// 		client.Snapshot(). = func() {
//
// 		}
// 		return client, nil
// 	}
//
// 	err := consulInstance.Backup()
//
// 	assert.NotNil(t, err)
// 	assert.Equal(t, err.Error(), "could not save snapshot")
// }
