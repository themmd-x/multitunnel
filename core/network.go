package core

import "net"

type NetworkListener interface {
    net.Listener
	ID() int64
	PongData([]byte)
}
