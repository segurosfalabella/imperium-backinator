package consul

import (
	"errors"
	"fmt"
	"regexp"

	"github.com/hashicorp/consul/api"
)

//API var
var API = api.NewClient

//Backuper struct
type Backuper struct {
	Endpoint string
	Token    string
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

	_, err := API(config)
	if err != nil {
		return errors.New("could not create new client")
	}

	fmt.Println(queryOptions)
	/*
		readCloser, _, err := client.Snapshot().Save(queryOptions)
		if err != nil {
			panic(err)
		}

		defer readCloser.Close()
		buf := new(bytes.Buffer)
		buf.ReadFrom(readCloser)
		newStr := buf.String()

		file, err := os.Create("backup.tgz")
		if err != nil {
			panic(err)
		}
		defer file.Close()
		file.WriteString(newStr)
	*/
	return nil
}

//ValidateEndpoint ...
func ValidateEndpoint(endpoint string) bool {
	expression, _ := regexp.Compile(`^(?:http(s)?:\/\/)?[\w.-]+(?:\.[\w\.-]+)+[\w\-\._~:/?#[\]@!\$&'\(\)\*\+,;=.]+$`)
	return expression.MatchString(endpoint)
}
