package v12140

import (
	"gophertunnel/v12140/minecraft"
	"gophertunnel/v12140/minecraft/protocol/login"
)

func Connect(host string) error {
	dialer := minecraft.Dialer{
	}

	conn, err := dialer.Dial("raknet", host)
	if err != nil {
		return err
	}
	defer conn.Close()

	if err := conn.DoSpawn(); err != nil {
		return err
	}
	for {}
	return nil
}
