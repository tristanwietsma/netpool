package netpool

import "testing"

func TestNewConnection(t *testing.T) {
	_, err := NewConnection("tcp", "127.0.0.1:9876")
	if err != nil {
		t.Fatal(err)
	}
}

func TestWrite(t *testing.T) {
	c, err := NewConnection("tcp", "127.0.0.1:9876")
	if err != nil {
		t.Fatal(err)
	}
	err = c.Write([]byte("ping"))
	if err != nil {
		t.Fatal(err)
	}
}

func TestWriteRead(t *testing.T) {
	c, err := NewConnection("tcp", "127.0.0.1:9876")
	if err != nil {
		t.Fatal(err)
	}
	err = c.Write([]byte("ping"))
	if err != nil {
		t.Fatal(err)
	}
	msg, err := c.Read()
	if err != nil {
		t.Fatal(err)
	}
	if string(msg) != "pong" {
		t.Fatal("did not match")
	}
}
