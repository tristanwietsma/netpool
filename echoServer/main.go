package main

import (
    "net"
)

func main() {
    ln, _ := net.Listen("tcp", "127.0.0.1:9876")
    defer ln.Close()

    for {
        conn, _ := ln.Accept()

        go func(c net.Conn) {
            buf := make([]byte, 1024)
            c.Read(buf)
            c.Write([]byte("pong"))
        }(conn)
    }

}
