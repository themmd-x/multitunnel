package main

import (
	"fmt"
	"github.com/TheMMD-X/multitunnel/minecraft"
	"github.com/TheMMD-X/multitunnel/minecraft/protocol/login"
)

func main() {
	address := "127.0.0.1:19132"

	dialer := minecraft.Dialer{
		Version: "1.21.50",
		IdentityData: login.IdentityData {
			DisplayName: "MultiTunnelFan",
		},
	}

	conn, err := dialer.Dial("raknet", address)
	if err != nil {
		panic(err)
	}
	fmt.Println("Connected")

	if err := conn.DoSpawn(); err != nil {
		panic(err)
	}
	fmt.Println("Spawned")
	select {}
}
