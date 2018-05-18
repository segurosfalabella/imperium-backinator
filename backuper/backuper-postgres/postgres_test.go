package postgres_test

import (
	"testing"

	postgres "github.com/segurosfalabella/imperium-backinator/backuper/backuper-postgres"
	"github.com/stretchr/testify/assert"
)

var port = 5432
var user = "root"
var password = "abc123"
var host = "127.0.0.1"

func TestShouldReturnErrorWhenHasNoHost(t *testing.T) {
	postgres := postgres.Backuper{
		Host:     "",
		Port:     port,
		User:     user,
		Password: password,
	}

	err := postgres.Backup()

	assert.NotNil(t, err)
}

func TestShouldReturnErrorWhenHasNoValidHost(t *testing.T) {
	postgres := postgres.Backuper{
		Host:     "Hola mundo",
		Port:     port,
		User:     user,
		Password: password,
	}

	err := postgres.Backup()

	assert.NotNil(t, err)
}

func TestShouldReturnErrorWhenHasNoPort(t *testing.T) {
	postgres := postgres.Backuper{
		Host:     host,
		User:     user,
		Password: password,
	}

	err := postgres.Backup()
	assert.NotNil(t, err)
}

func TestShouldReturnErrorWhenHasNoUser(t *testing.T) {
	postgres := postgres.Backuper{
		Host:     host,
		Port:     port,
		Password: password,
	}

	err := postgres.Backup()

	assert.NotNil(t, err)
}

func TestShouldReturnErrorWhenHasNoPassword(t *testing.T) {
	postgres := postgres.Backuper{
		Host: host,
		Port: port,
		User: user,
	}

	err := postgres.Backup()

	assert.NotNil(t, err)
}

func TestValidateMethodShouldReturnFalseWhenIpOrHostIsInvalid(t *testing.T) {
	var address = "hola mundo"

	result := postgres.ValidateHostFormat(address)

	assert.False(t, result)
}

func TestValidateMethodShouldReturnTrueWhenIpOrHostIsValid(t *testing.T) {
	var address = "127.0.0.1"

	result := postgres.ValidateHostFormat(address)

	assert.True(t, result)
}
