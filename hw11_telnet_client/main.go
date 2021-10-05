package main

import (
	"flag"
	"log"
	"os"
	"time"
)

var (
	timeOutFlag time.Duration
)

func init() {
	flag.DurationVar(&timeOutFlag, "timeout", time.Second*10, "connection timeout")
}

func main() {
	flag.Parse()
	args := flag.Args()
	if len(args) != 2 {
		log.Fatal("Enter a host and port number")
	}
	address := args[0] + ":" + args[1]
	telnet := NewTelnetClient(address, timeOutFlag, os.Stdin, os.Stdout)
	if err := telnet.Connect(); err != nil {
		log.Fatal(err)
	}
}
