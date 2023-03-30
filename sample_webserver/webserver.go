package main

import (
	"flag"
	"sample_webserver/server"
)

func main() {
	listenPort := flag.String("server-port", "18080", "Port at which Server Will be running")
	flag.Parse()
	server.Server(*listenPort)
}
