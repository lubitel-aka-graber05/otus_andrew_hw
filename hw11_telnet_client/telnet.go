package main

import (
	"bufio"
	"context"
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"sync"
	"time"
)

type TelnetClient interface {
	Connect() error
	io.Closer
	Send() error
	Receive() error
}

type Telnet struct{}

func (t *Telnet) Connect() error {
	connectTo()
}

func connectTo() (context.Context, net.Conn) {
	dialer := &net.Dialer{}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	conn, err := dialer.DialContext(ctx, "tcp", "127.0.0.1:3302")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	return ctx, conn
}

func (t *Telnet) Send() error {
	if err := writeMessage(connectTo()); err != nil {
		return err
	}
	return nil
}

func (t *Telnet) Receive() error {
	if err := readMessage(connectTo()); err != nil {
		return err
	}
	return nil
}

func readMessage(ctx context.Context, conn net.Conn) error {
	//buf:=bufio.NewScanner(conn)
	messageToOut := bufio.NewWriter(os.Stdout)
	for {
		select {
		case <-ctx.Done():
			break
		default:
			if _, err := messageToOut.ReadFrom(conn); err != nil {
				return err
			}
		}
	}
}

func writeMessage(ctx context.Context, conn net.Conn) error {
	messageToIn := bufio.NewWriter(conn)
	for {
		select {
		case <-ctx.Done():
			break
		default:
			if _, err := messageToIn.ReadFrom(os.Stdin); err != nil {
				return err
			}
		}
	}
}

func (t *Telnet) Close() error {
	return nil
}

func NewTelnetClient(address string, timeout time.Duration, in io.ReadCloser, out io.Writer) TelnetClient {
	// Place your code here.
	return &Telnet{}
}

// Place your code here.
// P.S. Author's solution takes no more than 50 lines.
