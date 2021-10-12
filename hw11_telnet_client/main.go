package main

import (
	"context"
	"flag"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	var timeOutFlag time.Duration
	flag.DurationVar(&timeOutFlag, "timeout", time.Second*10, "connection timeout")
	flag.Parse()
	args := flag.Args()
	log.Println(len(args))
	if len(args) != 2 {
		log.Fatal("Enter a host and port number")
	}

	telnet := NewTelnetClient(net.JoinHostPort(args[0], args[1]), timeOutFlag, os.Stdin, os.Stdout)
	if err := telnet.Connect(); err != nil {
		log.Fatal(err)
	}

	ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer cancel()

	go func() {
		if err := telnet.Send(); err != nil {
			log.Fatal(err)
		}
		log.Println("connection closed")
		cancel()
	}()
	go func() {
		if err := telnet.Receive(); err != nil {
			log.Fatal(err)
		}
	}()

	<-ctx.Done()
}
