package main

import (
	"fmt"
	"log"
	"net"
	"os"

	tcp "tcp/functions"
)

func main() {
	_, _ = os.Create("history.txt")

	//1 step launching the server
	port := tcp.Specify_port()
	listener, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("Error while listening to port %s : %s", port, err)
	}

	defer listener.Close()

	fmt.Printf("Server listen on port: %s\n", port)

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Printf("Error while establishing connection : %s", err)
			continue
		} else {
			go tcp.HandleClient(conn)

		}
	}
}
