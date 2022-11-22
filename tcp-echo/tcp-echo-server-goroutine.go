package main

import (
	"fmt"
	"net"
	"os"
	"runtime"
)

func main() {
	fmt.Println("> echo-server is activated")
	listen, err := net.Listen("tcp", "127.0.0.1:65456")
	if err != nil {
		fmt.Println("> listen failed and program terminated")
		return
	}
	defer listen.Close()
	for {
		socket, err := listen.Accept()
		if err != nil {
			fmt.Println("> accept failed and program terminated")
		}
		go echo01(socket)
	}
	fmt.Println("> ehco-server is de-activated")
}

func echo01(socket net.Conn) {
	fmt.Println("> client connected by", socket.LocalAddr().String())
	recvData := make([]byte, 1024)
	defer socket.Close()
	num := runtime.NumGoroutine()
	for {
		size, err := socket.Read(recvData)
		if err != nil {

		}

		if size > 0 {
			recvStr := string(recvData[:size])
			fmt.Printf("> echoed: %s by Goroutine-%d\n", recvStr, num)
			socket.Write([]byte(recvStr))
			if recvStr == "quit" {
				if runtime.NumGoroutine() == 2 {
					fmt.Println("> ehco-server is de-activated")
					os.Exit(0)
				}
				break
			}
		}
	}
}
