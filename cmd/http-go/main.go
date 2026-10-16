package main

import (
	"fmt"
	"log"
	"net"
)

func main() {
	socket, err := net.Listen("tcp", ":8000")
	if err != nil {
		log.Fatalln("Error in Starting Listener", err)
	}
	defer func() {
		fmt.Println("Connection Closed")
		socket.Close()
	}()
	fmt.Println("Application Listening on Port 8000")
	// we need to keep on listening to this otherwise it will close itself
	for {
		conn, err := socket.Accept()
		if err != nil {
			log.Fatalln("Error Accepting Connection", err)
		}
		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	// now we got our connection, lets see if it actually being called or not
	defer conn.Close()
	fmt.Println("Recieved one Connection!", conn)
}
