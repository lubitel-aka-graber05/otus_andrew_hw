package main

import (
	"bufio"
	"context"
	"fmt"
	"io"
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

type Telnet struct {
	addr    string
	timeOut time.Duration
	in      io.ReadCloser
	out     io.Writer
	wg      sync.WaitGroup
}

func (t *Telnet) Connect() error {
	dialer := &net.Dialer{}
	ctx, cancel := context.WithTimeout(context.Background(), t.timeOut)
	defer cancel()

	conn, err := dialer.DialContext(ctx, "tcp", t.addr)
	if err != nil {
		return fmt.Errorf("Error from Connect.DialContext: %w\n", err)
	}
	defer conn.Close()

	return nil
}

func (t *Telnet) Send() error {
	outCh := make(chan string)
	outMessage := bufio.NewWriter(t.out)
	ctx, cancel := context.WithTimeout(context.Background(), t.timeOut)
	defer cancel()
	t.wg.Add(1)
	go func() {
		defer t.wg.Done()
		for {
			select {
			case <-ctx.Done():
				break
			default:
				buf := bufio.NewScanner(os.Stdin)

				for buf.Scan() {
					outCh <- buf.Text()
				}
				if buf.Err() != nil {
					close(outCh)
				}
			}
		}
	}()
	t.wg.Wait()
	for s := range outCh {
		if _, err := outMessage.WriteString(s); err != nil {
			return fmt.Errorf("Error from Send.WriteString method: %w\n", err)
		}
	}
	return nil
}

func (t *Telnet) Receive() error {
	inCh := make(chan string)
	inMessage := bufio.NewWriter(os.Stdout)
	ctx, cancel := context.WithTimeout(context.Background(), t.timeOut)
	defer cancel()
	t.wg.Add(1)
	go func() {
		defer t.wg.Done()
		for {
			select {
			case <-ctx.Done():
				break
			default:
				buf := bufio.NewScanner(t.in)
				for buf.Scan() {
					inCh <- buf.Text()
				}
				if buf.Err() != nil {
					close(inCh)
					break
				}
			}
		}

	}()
	t.wg.Wait()
	for s := range inCh {
		if _, err := inMessage.WriteString(s); err != nil {
			return fmt.Errorf("Error from Recive.WriteString method: %w\n", err)
		}
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
