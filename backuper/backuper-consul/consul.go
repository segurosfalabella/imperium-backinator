package consul

import (
	"errors"
	"regexp"
)

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

	// queryOptions := &api.QueryOptions{
	// 	Token: cb.Token,
	// }
	//
	// config := &api.Config{
	// 	Address: cb.Endpoint,
	// }
	//
	// fmt
	return nil
}

//ValidateEndpoint ...
func ValidateEndpoint(endpoint string) bool {
	expression, _ := regexp.Compile(`^(?:http(s)?:\/\/)?[\w.-]+(?:\.[\w\.-]+)+[\w\-\._~:/?#[\]@!\$&'\(\)\*\+,;=.]+$`)
	return expression.MatchString(endpoint)
}
