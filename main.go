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

	// Create channels for incoming connections and server shutdown
	connect := make(chan net.Conn)
	off := make(chan bool)

	// 1 step launching the server
	port := tcp.Specify_port()
	listener, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("Error while listening to port %s : %s", port, err)
	}
	defer listener.Close()

	fmt.Printf("Server listen on port: %s\n", port)

	// Goroutine to handle incoming connections and signal shutdown
	go func() {
		for {
			select {
			case haveToOff := <-off:
				if haveToOff {
					return
				}
			default:
				conn, err := listener.Accept()
				if err != nil {
					log.Printf("Error while establishing connection: %s", err)
					continue
				}
				connect <- conn
			}
		}
	}()

	// Start a goroutine for each incoming connection
	for conn := range connect {
		go tcp.HandleClient(conn)
	}

	// Signal server shutdown and wait for goroutines to finish
	close(off)
}
