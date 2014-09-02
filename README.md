# netpool

[![Build Status](https://travis-ci.org/tristanwietsma/netpool.svg?branch=master)](https://travis-ci.org/tristanwietsma/netpool) [![GoDoc](https://godoc.org/github.com/tristanwietsma/netpool?status.svg)](http://godoc.org/github.com/tristanwietsma/netpool)

Connection Pool

## Install

```bash
go get github.com/tristanwietsma/netpool
```

## Usage

### Create a connection pool

```go
p, err := netpool.NewConnectionPool("<type>", "<server ip address>", <pool size>)
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
