package netpool

import (
	"log"
	"net"
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
	if err != nil {
		log.Printf("netpool: error on dial. %v\n", err)
		c.Close()
	}
	return &c, err
}

func (c *Connection) Write(msg []byte) error {
	_, err := c.conn.Write(msg)
	if err != nil {
		log.Printf("netpool: error on write. %v\n", err)
		c.Close()
	}
	return err
}

func (c *Connection) Read() ([]byte, error) {
	buf := make([]byte, 1024)
	nb, err := c.conn.Read(buf)
	if err != nil {
		log.Printf("netpool: error on read. %v\n", err)
		c.Close()
	}
	return buf[:nb], err
}

func (c *Connection) Close() error {
	err := c.conn.Close()
	c.IsClosed = true
	return err
}
