package netpool

import "testing"

func TestNewConnectionPool(t *testing.T) {
	_, err := NewConnectionPool("tcp", "127.0.0.1:9876", 3)
	if err != nil {
		t.Fail()
	}
}

func TestConnect(t *testing.T) {
	p, err := NewConnectionPool("tcp", "127.0.0.1:9876", 2)
	if err != nil {
		t.Fail()
	}

	// first (connect, release)
	c1, err := p.Connect()
	if err != nil {
		t.Fail()
	}

	if p.size != 1 {
		t.Fail()
	}

	p.Release(c1)
	if len(p.conn) != 1 {
		t.Fail()
	}

	// second (connect, release)
	c2, err := p.Connect()
	if err != nil {
		t.Fail()
	}

	if p.size != 1 {
		t.Fail()
	}

	p.Release(c2)
	if len(p.conn) != 1 {
		t.Fail()
	}

	// third (connect, hold)
	c3, err := p.Connect()
	if err != nil {
		t.Fail()
	}

	if p.size != 1 {
		t.Fail()
	}

	if len(p.conn) != 0 {
		t.Fail()
	}

	// fourth (connect, hold)
	c4, err := p.Connect()
	if err != nil {
		t.Fail()
	}

	if p.size != 2 {
		t.Fail()
	}

	if len(p.conn) != 0 {
		t.Fail()
	}

	// release one
	p.Release(c3)
	if len(p.conn) != 1 {
		t.Fail()
	}

	// release last one
	p.Release(c4)
	if len(p.conn) != 2 {
		t.Fail()
	}
}
