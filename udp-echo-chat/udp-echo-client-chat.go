package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	fmt.Println("> echo-client is activated")
	server_adr, err := net.ResolveUDPAddr("udp", "127.0.0.1:65456")
	local_addr, err := net.ResolveUDPAddr("udp", "127.0.0.1:")
	socket, err := net.DialUDP("udp", local_addr, server_adr)
	if err != nil {
		return
	}
	go udpClientRecv(socket)
	udpClientSend(socket)

	fmt.Println("> ehco-server is de-activated")
}

func udpClientRecv(socket *net.UDPConn) {
	for {
		recvData := make([]byte, 1024)
		size, _, err := socket.ReadFromUDP(recvData)
		if err != nil {
			fmt.Println("> ReadFrom failed and program terminated")
			return
		}
		if size > 0 {
			recvStr := string(recvData[:size])
			fmt.Printf("> received: %s\n", recvStr)
		}
	}
}

func udpClientSend(socket *net.UDPConn) {
	for {
		var sendMsg string
		fmt.Print("> ")
		fmt.Scanln(&sendMsg)
		socket.Write([]byte(sendMsg))
		if sendMsg == "quit" {
			fmt.Println("> ehco-client is de-activated")
			os.Exit(0)
		}
	}
}
