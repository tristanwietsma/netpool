package tcpool

import "net"

type Connection struct {
    conn net.Conn
}

// Returns a TCP connection to `address`.
func NewConnection(address string) (*Connection, error) {
    conn, err := net.Dial("tcp", address)
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
