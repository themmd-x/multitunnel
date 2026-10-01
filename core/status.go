package core

import (
	"net"
	"strconv"
	"sync"
	"time"

	raknet "github.com/TheMMD-X/multitunnel/libs/go-raknet/v1152"
)

type ForeignStatusProvider struct {
	addr string

	mu     sync.Mutex
	status ServerStatus

	once   sync.Once
	closed chan struct{}
}

func NewForeignStatusProvider(addr string) (*ForeignStatusProvider, error) {
	if _, err := net.ResolveUDPAddr("udp", addr); err != nil {
		return nil, err
	}
	f := &ForeignStatusProvider{addr: addr, closed: make(chan struct{})}
	go f.update()
	return f, nil
}

func (f *ForeignStatusProvider) ServerStatus(int, int) ServerStatus {
	f.mu.Lock()
	defer f.mu.Unlock()
	return f.status
}

func (f *ForeignStatusProvider) Close() error {
	f.once.Do(func() {
		close(f.closed)
	})
	return nil
}

func (f *ForeignStatusProvider) update() {
	ticker := time.NewTicker(time.Second)
	defer ticker.Stop()
	for {
		select {
		case <-ticker.C:
			data, err := raknet.Ping(f.addr)
			if err != nil {
				continue
			}
			f.mu.Lock()
			f.status = ParsePongData(data)
			f.mu.Unlock()
		case <-f.closed:
			return
		}
	}
}

func ParsePongData(pong []byte) ServerStatus {
	frag := SplitPong(string(pong))
	if len(frag) < 8 {
		return ServerStatus{ServerName: "Invalid pong data"}
	}
	online, err := strconv.Atoi(frag[4])
	if err != nil {
		return ServerStatus{ServerName: "Invalid player count"}
	}
	max, err := strconv.Atoi(frag[5])
	if err != nil {
		return ServerStatus{ServerName: "Invalid max player count"}
	}
	return ServerStatus{
		ServerName:    frag[1],
		ServerSubName: frag[7],
		PlayerCount:   online,
		MaxPlayers:    max,
	}
}

func SplitPong(s string) []string {
	var runes []rune
	var tokens []string
	inEscape := false
	for _, r := range s {
		switch {
		case r == '\\':
			inEscape = true
		case r == ';':
			tokens = append(tokens, string(runes))
			runes = runes[:0]
		case inEscape:
			inEscape = false
			fallthrough
		default:
			runes = append(runes, r)
		}
	}
	return append(tokens, string(runes))
}
