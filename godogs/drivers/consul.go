package drivers

import consul "github.com/segurosfalabella/imperium-backinator/backuper/backuper-consul"

// ConsulDriver ...
type ConsulDriver struct {
	Endpoint string
	Token    string
}

// Backup ...
func (cd *ConsulDriver) Backup() {
	consul := consul.Backuper{Endpoint: cd.Endpoint, Token: cd.Token}
	consul.Backup()
}
