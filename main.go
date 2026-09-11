package main

import (
	"fmt"
	"time"
	"strconv"
	"math/rand/v2"
	"gophertunnel/v12150/minecraft"
	"gophertunnel/v12150/minecraft/protocol/packet"
	"gophertunnel/v12150/minecraft/protocol/login"
)

func main() {

	address := "127.0.0.1:19132";

	dialer := minecraft.Dialer{}

	conn, err := dialer.Dial("raknet", &address)
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	if err := conn.DoSpawn(); err != nil {
		panic(err)
	}

	for {}
}
