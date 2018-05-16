package consul

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"os"
	"regexp"

	"github.com/hashicorp/consul/api"
)

//OsCreate var
var OsCreate = os.Create

//API var
var API = api.NewClient

//Backuper struct
type Backuper struct {
	Endpoint string
	Token    string
}

type CustomClient struct {
	client *api.Client
}

func (cc *CustomClient) Snapshot() *MySnapshot {
	return &MySnapshot{
		snapshot: &api.Snapshot{},
	}
}

//CustomClient interface
type CustomClientInterface interface {
	Snapshot() *MySnapshot
}

type MySnapshot struct {
	snapshot *api.Snapshot
}

func (ms *MySnapshot) Snapshot() *MySnapshot {
	return &MySnapshot{snapshot: &api.Snapshot{}}
}

func (ms *MySnapshot) Save(options *api.QueryOptions) (io.ReadCloser, *api.QueryMeta, error) {
	return ms.snapshot.Save(options)
}

func NewClient(config *api.Config) (CustomClientInterface, error) {
	client, err := api.NewClient(config)
	if err != nil {
		return nil, errors.New("could not upgrade new client method")
	}
	CustomClient := new(CustomClient)
	CustomClient.client = client
	return CustomClient, nil
}

// Backup functionality
func (cb *Backuper) Backup() error {
	if !ValidateEndpoint(cb.Endpoint) {
		return errors.New("endpoint not valid")
	}

	if cb.Token == "" {
		return errors.New("token is required")
	}

	queryOptions := &api.QueryOptions{
		Token: cb.Token,
	}

	config := &api.Config{
		Address: cb.Endpoint,
	}

	client, err := API(config)
	if err != nil {
		return errors.New("could not create new client")
	}

	readCloser, _, err := client.Snapshot().Save(queryOptions)
	if err != nil {
		fmt.Println(err.Error())
	}
	defer readCloser.Close()

	buf := new(bytes.Buffer)
	buf.ReadFrom(readCloser)
	newStr := buf.String()

	file, err := OsCreate("backup.tgz")
	if err != nil {
		return errors.New("could not create backup file")
	}

	defer file.Close()
	file.WriteString(newStr)

	return nil
}

//SaveSnapshot method
func SaveSnapshot(client CustomClient, options *api.QueryOptions) io.ReadCloser {
	readCloser, _, err := client.Snapshot().Save(options)
	if err != nil {
		return nil
	}
	return readCloser
}

//ValidateEndpoint ...
func ValidateEndpoint(endpoint string) bool {
	expression, _ := regexp.Compile(`^(?:http(s)?:\/\/)?[\w.-]+(?:\.[\w\.-]+)+[\w\-\._~:/?#[\]@!\$&'\(\)\*\+,;=.]+$`)
	return expression.MatchString(endpoint)
}
