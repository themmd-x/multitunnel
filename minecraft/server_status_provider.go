package minecraft

import (
	mtCore "github.com/TheMMD-X/multitunnel/core"
)

type (
	ServerStatus           = mtCore.ServerStatus
	ServerStatusProvider   = mtCore.ServerStatusProvider
	ListenerStatusProvider = mtCore.ListenerStatusProvider
	ForeignStatusProvider  = mtCore.ForeignStatusProvider
)

func NewStatusProvider(serverName, serverSubName string) ListenerStatusProvider {
	return mtCore.NewStatusProvider(serverName, serverSubName)
}

func NewForeignStatusProvider(addr string) (*ForeignStatusProvider, error) {
	return mtCore.NewForeignStatusProvider(addr)
}

func ParsePongData(pong []byte) ServerStatus {
	return mtCore.ParsePongData(pong)
}
