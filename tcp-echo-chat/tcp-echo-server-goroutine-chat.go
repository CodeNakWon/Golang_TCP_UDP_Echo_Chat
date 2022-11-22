package main

import (
	"fmt"
	"net"
	"os"
	"runtime"
)

var connections []net.Conn
var ch = make(chan string, 64)

func main() {

	fmt.Println("> echo-server is activated")
	listen, err := net.Listen("tcp", "127.0.0.1:65456")
	if err != nil {
		fmt.Println("> listen failed and program terminated")
		return
	}
	defer listen.Close()

	go sendMsgAll()

	for {
		socket, err := listen.Accept()
		if err != nil {
			fmt.Println("> accept failed and program terminated")
		}
		connections = append(connections, socket)
		go msgHandler(socket)
	}
	fmt.Println("> ehco-server is de-activated")
}

func sendMsgAll() {
	for {
		msg := <-ch
		for num := 0; num < len(connections); num++ {
			connections[num].Write([]byte(msg))
		}
	}
}

func msgHandler(socket net.Conn) {
	fmt.Println("> client connected by", socket.RemoteAddr().String())
	recvData := make([]byte, 1024)
	defer socket.Close()
	for {
		size, err := socket.Read(recvData)
		if err != nil {
		}
		if size > 0 {
			recvStr := string(recvData[:size])
			fmt.Printf("> received ( %s ) and echoed to %d clients\n", recvStr, runtime.NumGoroutine()-2)
			if recvStr == "quit" {
				if runtime.NumGoroutine() == 3 {
					fmt.Println("> ehco-server is de-activated")
					os.Exit(0)
				}
				break
			}
			ch <- recvStr
		}

	}
}
