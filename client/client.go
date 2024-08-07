package client

import "bufio"

type Client struct {
	host string
	port int
}

type Config struct {
	Host   string
	Client int
}

func (c *Client) dialer() {
	var reader *bufio.Reader
	var buf []byte
}
