package main

import (
	"fmt"
	"io"
	"net"
	"time"
)

type TelnetClient interface {
	Connect() error
	io.Closer
	Send() error
	Receive() error
}

type Telnet struct {
	address string
	timeOut time.Duration
	in      io.Reader
	out     io.Writer
	conn    net.Conn
}

func (t *Telnet) Connect() error {
	conn, err := net.DialTimeout("tcp", t.address, t.timeOut)
	if err != nil {
		return fmt.Errorf("connected: %w", err)
	}
	t.conn = conn
	return nil
}

func (t *Telnet) Send() error {
	if _, err := io.Copy(t.conn, t.in); err != nil {
		return fmt.Errorf("send: %w", err)
	}

	return nil
}

func (t *Telnet) Receive() error {
	if _, err := io.Copy(t.out, t.conn); err != nil {
		return fmt.Errorf("receive: %w", err)
	}

	return nil
}

func (t *Telnet) Close() (err error) {
	if t.conn != nil {
		if err = t.conn.Close(); err != nil {
			return fmt.Errorf("close: %w", err)
		}
	}
	return nil
}

func NewTelnetClient(address string, timeout time.Duration, in io.ReadCloser, out io.Writer) TelnetClient {
	return &Telnet{
		address: address,
		timeOut: timeout,
		in:      in,
		out:     out,
	}
}

// Place your code here.
// P.S. Author's solution takes no more than 50 lines.
