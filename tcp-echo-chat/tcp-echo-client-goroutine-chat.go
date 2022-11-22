package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	fmt.Println("> ehco-client is activated")
	socket, err := net.Dial("tcp", "127.0.0.1:65456")

	if err != nil {
		fmt.Println("> connect failed and program terminated")
	}
	defer socket.Close()

	go recvHandler(socket)
	sendMsgToServer(socket)
}

func sendMsgToServer(socket net.Conn) {
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

func recvHandler(socket net.Conn) {
	recvData := make([]byte, 1024)
	for {
		size, err := socket.Read(recvData)
		if err != nil {
			return
		}
		if size > 0 {
			recvMsg := string(recvData[:size])
			fmt.Println("> received:", recvMsg)
		}
	}
}
