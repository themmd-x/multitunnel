package minecraft

import (
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"
	"io"
	"net"
	"time"

	"github.com/TheMMD-X/multitunnel/versions"
)

const (
	batchHeader              = 0xfe
	idRequestNetworkSettings = 193
	packetIDMask             = 0x3ff
	detectTimeout            = 5 * time.Second
	detectBufferSize         = 1024
	fallbackBufferSize       = 1024 * 1024 * 3
)

type replayConn struct {
	net.Conn
	first []byte
	buf   []byte
}

func (c *replayConn) ReadPacket() ([]byte, error) {
	if c.first != nil {
		pk := c.first
		c.first = nil
		return pk, nil
	}
	if pr, ok := c.Conn.(interface{ ReadPacket() ([]byte, error) }); ok {
		return pr.ReadPacket()
	}
	if c.buf == nil {
		c.buf = make([]byte, fallbackBufferSize)
	}
	n, err := c.Conn.Read(c.buf)
	if err != nil {
		return nil, err
	}
	return c.buf[:n], nil
}

func (c *replayConn) Read(b []byte) (int, error) {
	if c.first != nil {
		if len(b) < len(c.first) {
			return 0, io.ErrShortBuffer
		}
		n := copy(b, c.first)
		c.first = nil
		return n, nil
	}
	return c.Conn.Read(b)
}

func (c *replayConn) Latency() time.Duration {
	if l, ok := c.Conn.(interface{ Latency() time.Duration }); ok {
		return l.Latency()
	}
	return 0
}

func detectVersion(conn net.Conn) (string, net.Conn, error) {
	type result struct {
		data []byte
		err  error
	}
	ch := make(chan result, 1)
	go func() {
		data, err := readFirstPacket(conn)
		ch <- result{data, err}
	}()

	var first []byte
	select {
	case r := <-ch:
		if r.err != nil {
			return "", nil, fmt.Errorf("read first packet: %w", r.err)
		}
		first = r.data
	case <-time.After(detectTimeout):
		return "", nil, errors.New("timed out waiting for the first packet")
	}

	protocol, err := parseClientProtocol(first)
	if err != nil {
		return "", nil, err
	}
	candidates := versions.ByProtocol[protocol]
	if len(candidates) == 0 {
		return "", nil, fmt.Errorf("unsupported client protocol %d", protocol)
	}
	return candidates[len(candidates)-1].Minecraft, &replayConn{Conn: conn, first: first}, nil
}

func readFirstPacket(conn net.Conn) ([]byte, error) {
	if pr, ok := conn.(interface{ ReadPacket() ([]byte, error) }); ok {
		return pr.ReadPacket()
	}
	buf := make([]byte, detectBufferSize)
	n, err := conn.Read(buf)
	if err != nil {
		return nil, err
	}
	return buf[:n], nil
}

func parseClientProtocol(data []byte) (int, error) {
	if len(data) == 0 || data[0] != batchHeader {
		return 0, errors.New("first packet is not a game packet batch")
	}
	r := bytes.NewReader(data[1:])
	length, err := binary.ReadUvarint(r)
	if err != nil {
		return 0, fmt.Errorf("read packet length: %w", err)
	}
	if length == 0 || length > uint64(r.Len()) {
		return 0, errors.New("invalid packet length")
	}
	header, err := binary.ReadUvarint(r)
	if err != nil {
		return 0, fmt.Errorf("read packet header: %w", err)
	}
	if id := header & packetIDMask; id != idRequestNetworkSettings {
		return 0, fmt.Errorf("first packet has id %d, expected RequestNetworkSettings", id)
	}
	var protocol [4]byte
	if _, err := io.ReadFull(r, protocol[:]); err != nil {
		return 0, fmt.Errorf("read client protocol: %w", err)
	}
	return int(int32(binary.BigEndian.Uint32(protocol[:]))), nil
}
