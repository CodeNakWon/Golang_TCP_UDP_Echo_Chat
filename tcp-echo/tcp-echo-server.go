package main

import (
	"fmt"
	"net"
)

func main() {
	fmt.Println("> echo-server is activated")
	listen, err := net.Listen("tcp", "127.0.0.1:65456")
	if err != nil {
		fmt.Println("> listen failed and program terminated")
		return
	}
	defer listen.Close()

	socket, err := listen.Accept()
	if err != nil {
		fmt.Println("> accept failed and program terminated")
	}
	echo00(socket)
	fmt.Println("> ehco-server is de-activated")
}

func echo00(socket net.Conn) {
	fmt.Println("> client connected by", socket.RemoteAddr().String())
	recvData := make([]byte, 1024)
	defer socket.Close()

	for {
		size, err := socket.Read(recvData)
		if err != nil {

		}

		if size > 0 {
			recvStr := string(recvData[:size])
			fmt.Println("> echoed:", recvStr)
			socket.Write([]byte(recvStr))
			if recvStr == "quit" {
				break
			}
		}

	}
}
