package main

import (
	"flag"

	consul "github.com/segurosfalabella/imperium-backinator/backuper/backuper-consul"
)

var source = flag.String("source", "consul", "Source server to backup")
var endpoint = flag.String("endpoint", "", "Consul API endpoint")
var token = flag.String("token", "", "Consul admin token")

func main() {
	flag.Parse()
	if *source == "consul" {
		consul := consul.Backuper{Endpoint: *endpoint, Token: *token}
		consul.Backup()
	}

	if *source == "postgres" {
	}
}
