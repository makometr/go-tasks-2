package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"time"
)

type TelnetClient struct {
	address    string
	port       int
	timeout    time.Duration
	conn       net.Conn
	connReader *bufio.Reader
}

func NewTelnetClient(address string, port int, timeout time.Duration) *TelnetClient {
	res := &TelnetClient{
		address: address,
		port:    port,
		timeout: timeout,
	}

	return res
}

func (tc *TelnetClient) Connect() error {
	conn, err := net.DialTimeout("tcp", tc.address, tc.timeout)
	if err != nil {
		return fmt.Errorf("%s : %w", "error dial to host", err)
	}
	tc.conn = conn
	tc.connReader = bufio.NewReader(tc.conn)
	fmt.Fprintf(os.Stderr, "%s %s\n", "err: closed by peer", tc.address)
	return nil
}

func (tc *TelnetClient) Close() error {
	return tc.conn.Close()
}

func main() {
	client := NewTelnetClient("/", 8080, time.Second*3)
	err := client.Connect()
	if err != nil {
		fmt.Println(err)
		return
	}
	client.Close()
}
