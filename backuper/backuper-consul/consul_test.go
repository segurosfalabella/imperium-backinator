package consul_test

import (
	"errors"
	"io"
	"net/http"
	"os"
	"testing"

	consul "github.com/segurosfalabella/imperium-backinator/backuper/backuper-consul"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var badEndpoint = "hola mundo"
var endpoint = "http://www.google.com"
var badToken = "demo-token"

type mockClient struct {
	mock.Mock
}

func (client *mockClient) Do(request *http.Request) (*http.Response, error) {
	args := client.Called(request)
	return args.Get(0).(*http.Response), args.Error(1)
}

func TestShouldReturnErrorWhenCanNotValidateEndpoint(t *testing.T) {
	consul := consul.Backuper{
		Endpoint: badEndpoint,
	}

	err := consul.Backup()

	assert.NotNil(t, err)
}

func TestShouldReturnErrorWhenTokenIsNotDefined(t *testing.T) {
	consul := &consul.Backuper{
		Token:    "",
		Endpoint: endpoint,
	}

	err := consul.Backup()

	assert.NotNil(t, err)
}

func TestShouldValidateMethodReturnTrueWhenEndpointIsValid(t *testing.T) {
	state := consul.ValidateEndpoint(endpoint)

	assert.True(t, state)
}

func TestShoudlValidateMethodReturnFalseWhenEndpointIsInvalid(t *testing.T) {
	state := consul.ValidateEndpoint(badEndpoint)

	assert.False(t, state)
}

func TestShouldReturnErrorWhenNewHttpRequestFail(t *testing.T) {
	consulInstance := consul.Backuper{
		Endpoint: endpoint,
		Token:    badToken,
	}
	var oldRequest = consul.HttpRequest
	defer func() {
		consul.HttpRequest = oldRequest
	}()
	consul.HttpRequest = func(method string, url string, body io.Reader) (*http.Request, error) {
		return nil, errors.New("could not create new request")
	}

	err := consulInstance.Backup()

	assert.NotNil(t, err)
}

func TestShouldReturnErrorWhenDoRequestReturnError(t *testing.T) {
	consulInstance := consul.Backuper{
		Endpoint: endpoint,
		Token:    badToken,
	}
	var oldRequest = consul.DoRequest
	defer func() {
		consul.DoRequest = oldRequest
	}()
	consul.DoRequest = func(client consul.ClientInterface, request *http.Request) (*http.Response, error) {
		return new(http.Response), errors.New("could not execute request")
	}
	err := consulInstance.Backup()

	assert.NotNil(t, err)
}

func TestShouldReturnErrorNotNilWhenDoFuncFail(t *testing.T) {
	client := new(mockClient)
	client.On("Do", new(http.Request)).Return(new(http.Response), errors.New("could not execute action"))
	response, err := consul.DoRequest(client, new(http.Request))

	assert.NotNil(t, err)
	assert.NotNil(t, response)
}

func TestShouldReturnErrorNotNilWhenOsCreateFunctionFail(t *testing.T) {
	consulInstance := consul.Backuper{
		Endpoint: endpoint,
		Token:    badToken,
	}
	var oldOsCreate = consul.OsCreate
	defer func() {
		consul.OsCreate = oldOsCreate
	}()
	consul.OsCreate = func(name string) (*os.File, error) {
		return nil, errors.New("could not create backup file")
	}

	err := consulInstance.Backup()

	assert.NotNil(t, err)
}
