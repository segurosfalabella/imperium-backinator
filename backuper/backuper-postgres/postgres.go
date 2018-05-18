package postgres

import (
	"errors"
	"fmt"
	"regexp"

	postgresBinary "github.com/segurosfalabella/imperium-backinator/binary-postgres"
)

//Backuper struct
type Backuper struct {
	Host     string
	Port     int
	User     string
	Password string
}

func init() {
	postgresBinary.RestoreAsset()
}

//Backup method
func (pb *Backuper) Backup() error {
	if pb.Host == "" || !ValidateHostFormat(pb.Host) {
		return errors.New("host is required")
	}

	if pb.Port == 0 {
		return errors.New("port is required")
	}

	if pb.User == "" {
		return errors.New("user is required")
	}

	if pb.Password == "" {
		return errors.New("password is required")
	}

	return nil
}

//ValidateHostFormat method
func ValidateHostFormat(host string) bool {
	ipFormat := `\d{1,3}\.\d{1,3}\.\d{1,3}\.\d{1,3}`
	hostnameFormat := `^(?:http(s)?:\/\/)?[\w.-]+(?:\.[\w\.-]+)+[\w\-\._~:/?#[\]@!\$&'\(\)\*\+,;=.]+$`
	format := fmt.Sprintf(`(%s)|(%s)`, ipFormat, hostnameFormat)
	expression, _ := regexp.Compile(format)
	return expression.MatchString(host)
}
