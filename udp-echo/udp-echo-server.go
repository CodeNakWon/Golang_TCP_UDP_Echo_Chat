package main

import (
	"fmt"
	"net"
)

func main() {

	fmt.Println("> echo-server is activated")
	go udpEchoServer()
	for {
		var msg string
		fmt.Print("> ")
		fmt.Scanln(&msg)
		if msg == "quit" {
			fmt.Println("> stop procedure started")
			break
		}
	}
	fmt.Println("> ehco-server is de-activated")
}

func udpEchoServer() {
	socket, err := net.ListenPacket("udp", "127.0.0.1:65456")
	for {
		if err != nil {
			fmt.Println("> ListenPacket failed and program terminated", err)
			return
		}
		recvData := make([]byte, 1024)
		size, client, err := socket.ReadFrom(recvData)
		if err != nil {
			fmt.Println("> ReadFrom failed and program terminated")
			return
		}
		if size > 0 {
			recvStr := string(recvData[:size])
			fmt.Println("> echoed:", recvStr)
			socket.WriteTo([]byte(recvStr), client)
		}
	}
}
