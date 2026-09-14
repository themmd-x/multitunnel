package v12150

import (
	"github.com/TheMMD-X/multitunnel/libs/gophertunnel/v12150/minecraft"
)

func Connect(host string) (*minecraft.Conn, error) {
	dialer := minecraft.Dialer{}

	conn, err := dialer.Dial("raknet", host)
	if err != nil {
		return nil, err
	}

	return conn, nil
}
