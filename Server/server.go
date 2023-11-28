package main

import (
	"net"
)

func handleConnection(conn net.Conn) {
	defer conn.Close()

	buffer := make([]byte, 1024) //1024 bytes
}
