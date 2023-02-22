package main

import (
	"fmt"
	"net"
)

func main() {

	// Create a TCP listener on port 8080
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer listener.Close()

	fmt.Println("Listening on", listener.Addr())

	// Handle incoming connections
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Error accepting connection:", err)
			continue
		}

		fmt.Println("Connection from", conn.RemoteAddr())
		// create a map that will hold details of all the connection
		mp := make(map[string]net.Conn)

		// Handle connection in a separate goroutine

		go handleConnection(conn, mp)
	}

}

func handleConnection(con net.Conn, mp map[string]net.Conn) {
	con.Write([]byte("Hey there connected"))

}
