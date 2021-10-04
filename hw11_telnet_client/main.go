package main

import (
	"flag"
	"log"
	"os"
	"time"
)

var (
	host        string
	port        string
	timeOutFlag time.Duration
)

func init() {
	flag.StringVar(&host, "host", "127.0.0.1", "host address for connection")
	flag.StringVar(&port, "port", "4242", "port number")
	flag.DurationVar(&timeOutFlag, "timeout", time.Second*10, "connection timeout")
}

func main() {
	// Place your code here,
	// P.S. Do not rush to throw context down, think think if it is useful with blocking operation?
	flag.Parse()
	address := host + ":" + port
	telnet := NewTelnetClient(address, timeOutFlag, os.Stdin, os.Stdout)
	if err := telnet.Connect(); err != nil {
		log.Fatal(err)
	}
}
