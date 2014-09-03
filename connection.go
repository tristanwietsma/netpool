package netpool

import (
	"net"
	"time"
)

type Connection struct {
	conn     net.Conn
	IsClosed bool
}

// Returns a connection to `address`.
func NewConnection(network, address string) (*Connection, error) {
	conn, err := net.Dial(network, address)
	c := Connection{
		conn: conn,
	}
	if err == nil {
		c.IsClosed = true
	}
	return &c, err
}

func (c *Connection) Write(msg []byte) error {
	_, err := c.conn.Write(msg)
	if err != nil {
		c.IsClosed = true
	}
	return err
}

func (c *Connection) Read() ([]byte, error) {
	buf := make([]byte, 1024)
	nb, err := c.conn.Read(buf)
	if err != nil {
		c.IsClosed = true
	}
	return buf[:nb], err
}

func (c *Connection) Close() error {
	err := c.conn.Close()
	c.IsClosed = true
	return err
}
