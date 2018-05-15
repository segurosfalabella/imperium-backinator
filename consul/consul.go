package consul

import (
	"bytes"
	"os"

	"github.com/hashicorp/consul/api"
)

// Backup functionality
func Backup(endpoint string, token string) {
	queryOptions := &api.QueryOptions{
		Token: token,
	}

	config := &api.Config{
		Address: endpoint,
	}
	client, err := api.NewClient(config)
	if err != nil {
		panic(err)
	}
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
}
