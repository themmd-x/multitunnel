# MultiTunnel

Multi Version Minecraft Bedrock Protocol Library with an API similar to [GopherTunnel](github.com/Sandertv/gophertunnel)

## Features

- Packet level control
- Support client and server sides
- Runtime and compile time version selection
- Near equal API to gophertunnel
- Multi version servers

## GopherTunnel Compatibility

MultiTunnel intentionally follows the GopherTunnel API where possible.
For many applications, replacing:

github.com/sandertv/gophertunnel

with:

github.com/TheMMD-X/multitunnel
requires near no changes.

## Client example

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
}
```

## Supported versions

- 1.21.X
- 1.26.X

## Future updates

- Adding the missing pieces from specific structs

