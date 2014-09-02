package netpool

import "net"

type Connection struct {
	conn net.Conn
}

// Returns a connection to `address`.
func NewConnection(network, address string) (*Connection, error) {
	conn, err := net.Dial(network, address)
	c := Connection{
		conn: conn,
	}
	return &c, err
}

func (c *Connection) Write(msg []byte) error {
	_, err := c.conn.Write(msg)
	return err
}

func (c *Connection) Read() ([]byte, error) {
	buf := make([]byte, 1024)
	nb, err := c.conn.Read(buf)
	return buf[:nb], err
}
