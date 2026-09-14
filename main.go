package main

import (
	"fmt"
	"log"

	"github.com/TheMMD-X/multitunnel/minecraft"
	"github.com/TheMMD-X/multitunnel/query"
)

func main() {
	address := "127.0.0.1:19132"

	data, _ := query.Do(address)
	fmt.Println(data["version"])
	dialer := minecraft.Dialer{
		Version: "1.21.2",
	}

	conn, err := dialer.Dial("raknet", address)
	if err != nil {
		log.Fatal(err)
	}

	if err := conn.DoSpawn(); err != nil {
		log.Fatal(err)
	}
	for {}
}
