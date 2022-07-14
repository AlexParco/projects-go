package client

import (
	"bufio"
	"fmt"
	"net"
)

type Client struct {
	Nick   string
	Msg    chan string
	Addr   net.Addr
	Conn   net.Conn
	Reader *bufio.Reader
}

func New(conn net.Conn) *Client {
	reader := bufio.NewReader(conn)
	c := &Client{
		Nick:   "Anom",
		Msg:    make(chan string),
		Addr:   conn.RemoteAddr(),
		Conn:   conn,
		Reader: reader,
	}
	go c.Read()
	return c
}

func (c *Client) Read() {
	for {
		str, err := c.Reader.ReadString('\n')

		if err != nil {
			break
		}
		if str == "\n" {
			continue
		}
		c.Msg <- str
	}
}

func (c *Client) Write(message string) {
	c.Conn.Write([]byte(fmt.Sprintf("[%s]: %s", c.Nick, message)))
}
