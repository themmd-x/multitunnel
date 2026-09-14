package v12130

import (
	"github.com/TheMMD-X/multitunnel/libs/gophertunnel/v12130/minecraft"
)

func Connect(host string) (*minecraft.Conn, error) {
	dialer := minecraft.Dialer{}

	conn, err := dialer.Dial("raknet", host)
	if err != nil {
		return nil, err
	}

	return conn, nil
}
