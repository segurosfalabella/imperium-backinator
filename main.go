package main

import (
	"bytes"
	"os"

	"github.com/hashicorp/consul/api"
)

//QueryOptions type
// type QueryOptions struct {
// //Datacenter string
// Token string
// }

//Token var
var token = "7de16b0a-04a0-6bd9-f323-74cb6f8c330a"
var address = "http://consul-v2.tools.segurosfalabella.cloud"

func main() {
	queryOptions := &api.QueryOptions{
		Token: token,
	}

	config := &api.Config{
		Address: address,
	}
	client, err := api.NewClient(config)
	if err != nil {
		panic(err)
	}
	readCloser, q, err := client.Snapshot().Save(queryOptions)
	if err != nil {
		panic(err)
	}

	defer readCloser.Close()
	buf := new(bytes.Buffer)
	buf.ReadFrom(readCloser)
	newStr := buf.String()

	file, err := os.Create("demo.tgz")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	file.WriteString(newStr)

}
