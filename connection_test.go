package tcpool

import "testing"

func TestNewConnection(t *testing.T) {
    _, err := NewConnection("127.0.0.1:9876")
    if err != nil {
        t.Fail()
    }
}

func TestWrite(t *testing.T) {
    c, _ := NewConnection("127.0.0.1:9876")
    err := c.Write([]byte("ping"))
    if err != nil {
        t.Fail()
    }
}

func TestWriteRead(t *testing.T) {
    c, _ := NewConnection("127.0.0.1:9876")
    err := c.Write([]byte("ping"))
    if err != nil {
        t.Fail()
    }
    msg, err := c.Read()
    if err != nil {
        t.Fail()
    }
    if string(msg) != "pong" {
        t.Fail()
    }
}
