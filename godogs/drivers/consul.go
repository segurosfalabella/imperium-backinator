package drivers

import "github.com/segurosfalabella/imperium-backinator/consul"

// ConsulDriver ...
type ConsulDriver struct {
	Endpoint string
	Token    string
}

// Backuper ...
func (cd *ConsulDriver) Backuper() {
	consul.Backup(cd.Endpoint, cd.Token)

}
