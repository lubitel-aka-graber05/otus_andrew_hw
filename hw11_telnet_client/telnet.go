package main

import (
	"bufio"
	"context"
	"fmt"
	"io"
	"net"
	"os"
	"time"
)

type TelnetClient interface {
	Connect() error
	io.Closer
	Send() error
	Receive() error
}

type Telnet struct {
	addr    string
	timeOut time.Duration
	in      io.ReadCloser
	out     io.Writer
	ctx     context.Context
}

func (t *Telnet) Connect() error {
	dialer := &net.Dialer{}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	conn, err := dialer.DialContext(ctx, "tcp", t.addr)
	if err != nil {
		return fmt.Errorf("Error from DialContext: %w\n", err)
	}
	defer conn.Close()

	return nil
}

func (t *Telnet) Send() error {
	outCh := make(chan string)
	outMessage := bufio.NewWriter(t.out)
	go func() {
		buf := bufio.NewScanner(os.Stdin)

		for buf.Scan() {
			outCh <- buf.Text()
		}
		if buf.Err() != nil {
			close(outCh)
		}
	}()
	for s := range outCh {
		if _, err := outMessage.WriteString(s); err != nil {
			return fmt.Errorf("Error from WriteString method: %w\n", err)
		}
	}
	return nil
}

func (t *Telnet) Receive() error {
	//inCh:=make(chan string)

	buf := bufio.NewReader(t.in)
	if _, err := buf.WriteTo(os.Stdout); err != nil {
		return fmt.Errorf("Error from Recieve: %w\n", err)
	}
	return nil
}

func (t *Telnet) Close() error {
	return nil
}

func NewTelnetClient(address string, timeout time.Duration, in io.ReadCloser, out io.Writer) TelnetClient {
	// Place your code here.
	return &Telnet{
		addr:    address,
		timeOut: timeout,
		in:      in,
		out:     out,
	}
}

// Place your code here.
// P.S. Author's solution takes no more than 50 lines.
