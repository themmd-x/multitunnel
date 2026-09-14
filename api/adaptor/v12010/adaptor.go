package v12010

import (
	"gophertunnel/v12010/minecraft"
)

func Connect(host string) error {
	dialer := minecraft.Dialer{}

	conn, err := dialer.Dial("raknet", host)
	if err != nil {
		return err
	}
	defer conn.Close()

	if err := conn.DoSpawn(); err != nil {
		return err
	}
	return nil
}
