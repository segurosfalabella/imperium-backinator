package postgres_test

import (
	"testing"

	postgres "github.com/segurosfalabella/imperium-backinator/backuper/backuper-postgres"
)

func TestShouldReturnErrorWhenBackuperHasNoParameters(t *testing.T) {
	postgres := postgres.Backuper{
		Host:     "",
		Port:     "",
		User:     "",
		Password: "",
	}
}
