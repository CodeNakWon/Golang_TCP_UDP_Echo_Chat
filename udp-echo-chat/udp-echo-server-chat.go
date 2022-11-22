package main

import (
	"fmt"
	"net"
)

var m map[string]net.Addr

func main() {
	m = make(map[string]net.Addr)

	fmt.Println("> echo-server is activated")
	udpRecv()

	fmt.Println("> ehco-server is de-activated")
}

func udpRecv() {
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

			if recvStr == "#REG" {
				fmt.Println(client.String())
				m[client.String()] = client
				fmt.Printf("> client registerd ( %s )\n", client.String())
			} else if recvStr == "#DEREG" {
				delete(m, client.String())
			}
			fmt.Printf("> received ( %s ) and echoed to %d clients\n", recvStr, len(m))
			for _, val := range m {
				socket.WriteTo([]byte(recvStr), val)
			}

		}
	}
}
