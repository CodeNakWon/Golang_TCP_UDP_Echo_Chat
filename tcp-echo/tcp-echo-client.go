package main

import (
	"fmt"
	"net"
)

func main() {
	fmt.Println("> ehco-client is activated")
	socket, err := net.Dial("tcp", "127.0.0.1:65456")
	if err != nil {
		fmt.Println("> connect failed and program terminated")
	}
	defer socket.Close()

	recvData := make([]byte, 1024)
	for {
		var sendMsg string
		fmt.Print("> ")
		fmt.Scanln(&sendMsg)
		socket.Write([]byte(sendMsg))
		size, err := socket.Read(recvData)
		if err != nil {
			return
		}
		if size > 0 {
			recvMsg := string(recvData[:size])
			fmt.Println("> received:", recvMsg)
			if recvMsg == "quit" {
				fmt.Println("> ehco-client is de-activated")
				return
			}
		}
	}
}
