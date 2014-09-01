# tcpool

[![Build Status](https://travis-ci.org/tristanwietsma/tcpool.svg?branch=master)](https://travis-ci.org/tristanwietsma/tcpool) [![GoDoc](https://godoc.org/github.com/tristanwietsma/tcpool?status.svg)](http://godoc.org/github.com/tristanwietsma/tcpool)

TCP Connection Pool

## Install

```bash
go get github.com/tristanwietsma/tcpool
```

## Usage

### Create a connection pool

```go
p, err := tcpool.NewConnectionPool("<server ip address>", <pool size>)
```

### Get a connection from the pool

```go
c, err := p.Connect()
```

### Use a connection

Connections implement `Write` and `Read`.

```go
err := c.Write([]byte("ping"))
msg, err := c.Read()
```

### Return a connection to the pool

```go
p.Release(c)
```

## Testing

Build and start the `echoServer`, then `go test` as normal.
