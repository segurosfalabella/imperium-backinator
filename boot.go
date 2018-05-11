package main

import (
    "flag"
    "fmt"
    maria "github.com/segurosfalabella/imperium-backinator/binary-maria"
    mysql "github.com/segurosfalabella/imperium-backinator/binary-mysql"
)

var endpoint = flag.String("endpoint", "http://consul-v2.tools.segurosfalabella.cloud", "Consul API endpoint")
var token = flag.String("token", "", "Consul admin token")

func main() {
    fmt.Println(maria.AssetNames())
    fmt.Println(mysql.AssetNames())
    // flag.Parse()
    // app.Backup(*endpoint, *token)
}
