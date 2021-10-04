package main

import (
	"bufio"
	"context"
	"io"
	"log"
	"net"
	"sync"
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
	wg      sync.WaitGroup
	closer  io.Closer
	f       func() error
}

func (t *Telnet) Connect() error {
	dial := &net.Dialer{}
	ctx, cancel := context.WithTimeout(context.Background(), t.timeOut)
	defer cancel()
	var err error
	t.conn, err = dial.DialContext(ctx, "tcp", t.address)
	if err != nil {
		return err
	}

	t.wg.Add(1)
	go func() {
		defer t.wg.Done()
		for {
			select {
			case <-ctx.Done():
				break
			default:
				if err = t.Send(); err != nil {
					log.Println(err)
				}
			}
		}
	}()

	t.wg.Add(1)
	go func() {
		defer t.wg.Done()
		for {
			select {
			case <-ctx.Done():
				break
			default:
				if err = t.Receive(); err != nil {
					log.Println(err)
				}
			}
		}

	}()
	t.wg.Wait()

	if err = t.conn.Close(); err != nil {
		return err
	}

	return nil
}

func (t *Telnet) Send() error {
	defer t.wg.Done()

	message := bufio.NewScanner(t.in)
	for message.Scan() {
		if _, err := t.conn.Write(message.Bytes()); err != nil {
			return err
		}
		if _, err := t.conn.Write([]byte("\n")); err != nil {
			return err
		}
	}

	return nil
}

func (t *Telnet) Receive() error {
	defer t.wg.Done()

	for {
		writer := bufio.NewReader(t.conn)
		if _, err := writer.WriteTo(t.out); err != nil {
			return err
		}
	}
}

func (t *Telnet) Close() (err error) {
	if t.closer == nil {
		return nil
	}
	defer func() {
		err = t.closer.Close()
	}()
	return t.f()
}

func NewTelnetClient(address string, timeout time.Duration, in io.ReadCloser, out io.Writer) TelnetClient {
	// Place your code here.
	return &Telnet{
		address: address,
		timeOut: timeout,
		in:      in,
		out:     out,
	}
}

// Place your code here.
// P.S. Author's solution takes no more than 50 lines.
