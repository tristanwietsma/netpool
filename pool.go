package netpool

import "sync"

type ConnectionPool struct {
	sync.Mutex
	size    uint
	maxSize uint
	network string
	address string
	conn    chan *Connection
}

func NewConnectionPool(network, address string, size uint) (*ConnectionPool, error) {
	p := ConnectionPool{}
	p.size = 1
	p.maxSize = size
	p.network = network
	p.address = address
	p.conn = make(chan *Connection, size)
	conn, err := NewConnection(network, address)
	p.conn <- conn
	return &p, err
}

func (p *ConnectionPool) Connect() (*Connection, error) {

	// grab (or wait for) an existing connection
	if len(p.conn) > 0 || p.size == p.maxSize {
		return <-p.conn, nil
	}

	// create a new connection
	p.Lock()
	conn, err := NewConnection(p.network, p.address)
	if err == nil {
		p.size++
	}
	p.Unlock()

	return conn, err
}

func (p *ConnectionPool) Release(c *Connection) {
	p.conn <- c
}
