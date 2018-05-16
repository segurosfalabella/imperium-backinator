package consul

import (
	"bytes"
	"errors"
	"net/http"
	"os"
	"regexp"
)

//OsCreate var
var OsCreate = os.Create

//HttpRequest var
var HttpRequest = http.NewRequest

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

	request, err := HttpRequest("GET", cb.Endpoint+"/v1/snapshot", nil)

	if err != nil {
		return errors.New("could not create new request")
	}

	request.Header.Set("X-Consul-Token", cb.Token)
	client := &http.Client{}
	response, err := DoRequest(client, request)

	if err != nil {
		return err
	}

	buf := new(bytes.Buffer)
	buf.ReadFrom(response.Body)
	newStr := buf.String()

	file, err := OsCreate("backup.tgz")
	defer file.Close()
	if err != nil {
		return errors.New("could not create backup file")
	}
	file.WriteString(newStr)
	return nil
}

//ClientInterface interface
type ClientInterface interface {
	Do(request *http.Request) (*http.Response, error)
}

//DoRequest function
var DoRequest = func(client ClientInterface, request *http.Request) (*http.Response, error) {
	return client.Do(request)
}

//ValidateEndpoint ...
func ValidateEndpoint(endpoint string) bool {
	expression, _ := regexp.Compile(`^(?:http(s)?:\/\/)?[\w.-]+(?:\.[\w\.-]+)+[\w\-\._~:/?#[\]@!\$&'\(\)\*\+,;=.]+$`)
	return expression.MatchString(endpoint)
}
