package test

import (
	"os"
	"path/filepath"

	"github.com/DATA-DOG/godog"
	"github.com/segurosfalabella/imperium-backinator/godogs/drivers"
)

var endpoint = "http://localhost:7700"
var consulDriver drivers.ConsulDriver

func aSourceEqualsToConsulAndEndpointAndToken(source, endpoint, token string) error {
	consulDriver = drivers.ConsulDriver{
		Endpoint: endpoint,
		Token:    token,
	}
	return nil
}

func consulBackuperIsExecuted() error {
	consulDriver.Backuper()
	return nil
}

func aMustBeSaveInALocalDirectory(arg1 string) error {
	path := "../backup.tgz"
	err := filepath.Walk(path, func(path string, info os.FileInfo, err error) error {

		if err != nil {
			return err
		}

		if info.Name() != "backup.tgz" {
			return godog.ErrPending
		}

		return nil
	})

	if err != nil {
		return godog.ErrPending
	}

	return nil
}

func aSourceAndAndAndAndAnd(arg1, arg2, arg3, arg4, arg5, arg6 string) error {
	return godog.ErrPending
}

func postgresOptionIsExecuted() error {
	return godog.ErrPending
}

func FeatureContext(s *godog.Suite) {
	// Consul
	s.Step(`^a source "([^"]*)" equals to consul and endpoint "([^"]*)" and token "([^"]*)"$`, aSourceEqualsToConsulAndEndpointAndToken)
	s.Step(`^consul backuper is executed$`, consulBackuperIsExecuted)

	// Postgres
	s.Step(`^a source "([^"]*)" and  "([^"]*)" and "([^"]*)" and "([^"]*)" and "([^"]*)" and "([^"]*)"$`, aSourceAndAndAndAndAnd)
	s.Step(`^postgres option is executed$`, postgresOptionIsExecuted)

	s.Step(`^a "([^"]*)" must be save in a local directory$`, aMustBeSaveInALocalDirectory)

}
