package main

import (
	"flag"

	"github.com/segurosfalabella/imperium-backinator/app"
)

var endpoint = flag.String("endpoint", "http://consul-v2.tools.segurosfalabella.cloud", "Consul API endpoint")
var token = flag.String("token", "", "Consul admin token")

func main() {
	flag.Parse()
	app.Backup(*endpoint, *token)
}
