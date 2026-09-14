# MultiTunnel

> Under development, Possible bugs, Internal modules may change

Multi Version [GopherTunnel](https://github.com/Sandertv/gophertunnel) library collection with an API similar to gophertunnel

## Features

- Support client based tasks
- Runtime and compile time version selection
- Very similar API to gophertunnel

## installation

```sh
go get github.com/TheMMD-X/multitunnel
```

## example

```go
package main

import (
        "github.com/TheMMD-X/multitunnel/minecraft"
)

func main() {
        address := "127.0.0.1:19132"

        dialer := minecraft.Dialer{
                Version: "1.21.2",
        }

        conn, err := dialer.Dial("raknet", address)
        if err != nil {
                panic(err)
        }

        if err := conn.DoSpawn(); err != nil {
                panic(err)
        }
}```

## Supported versions

- 1.21.X
- 1.26.X

## Future updates

- Server side implementation
