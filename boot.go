package main

import (
	"flag"

	"github.com/segurosfalabella/imperium-backinator/consul"
)

var source = flag.String("source", "consul", "Source server to backup")
var endpoint = flag.String("endpoint", "", "Consul API endpoint")
var token = flag.String("token", "", "Consul admin token")

func main() {
	flag.Parse()
	if *source == "consul" {
		consul.Backup(*endpoint, *token)
	}

	if *source == "postgres" {
		consul.Backup(*endpoint, *token)
	}
}
