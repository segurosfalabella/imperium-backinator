package test

import (
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/DATA-DOG/godog"
	"github.com/segurosfalabella/imperium-backinator-consul/app"
)

var endpoint = "http://localhost:7700"

// StartServer for faking a consul server
func StartServer() {
	http.HandleFunc("/v1/snapshot", snapshot)
	go http.ListenAndServe(":7700", nil)
}

func snapshot(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hola", nil)
}

func anEndpoint() error {
	StartServer()
	return nil
}

func iCallGetSnapshot() error {
	app.Backup(endpoint, "")
	return nil
}

func iGetSnapshotBackupFileNamed(fileName string) error {
	files, err := ioutil.ReadDir("./")
	if err != nil {
		log.Fatal(err)
	}
	for _, f := range files {
		fmt.Println(f.Name())
		if f.Name() == fileName {
			return nil
		}
	}
	return errors.New("there is no file named " + fileName)
}

func anSnapshot() error {
	return godog.ErrPending
}

func iCallCloudStorageEndpoint() error {
	return godog.ErrPending
}

func iPutBackupFileInCloudStorage() error {
	return godog.ErrPending
}

func FeatureContext(s *godog.Suite) {
	s.Step(`^an endpoint$`, anEndpoint)
	s.Step(`^I call get snapshot$`, iCallGetSnapshot)
	s.Step(`^I get snapshot backup file named "([^"]*)"$`, iGetSnapshotBackupFileNamed)
	s.Step(`^an snapshot$`, anSnapshot)
	s.Step(`^I call cloud storage endpoint$`, iCallCloudStorageEndpoint)
	s.Step(`^I put backup file in cloud storage$`, iPutBackupFileInCloudStorage)
}
