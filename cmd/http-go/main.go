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
			fmt.Println("Error Accepting Connection", err)
		}
		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	// now we got our connection, lets see if it actually being called or not
	defer func() {
		fmt.Println("Connection Closed")
		conn.Close()
	}()
	fmt.Println("Recieved one Connection!", conn)
	// now we also need to accept if what the connection might be sending us
	// let's read in chunk sizes and then stitch them together and print them then
	// for now, lets just read 10 chars, i.e 80 bytes, I wanna see how it reacts.
	// buffer := make([]byte, 1024)
	// for {
	// 	n, err := conn.Read(buffer)
	// 	if errorhandling.HandleIoError(err) {
	// 		return
	// 	}
	// 	raw_data := buffer[:n] // we will only read till we recieved
	// 	fmt.Println("This is the data we recieved\n", raw_data)
	// }
	// lets try sending some data instead of reading now
	text_to_send := "lol yo got this from my message, can't beleive it"
	bytes_to_send := []byte(text_to_send)
	n, err := conn.Write(bytes_to_send)
	if err != nil {
		fmt.Println("Error Sending Message", err)
	} else {
		fmt.Println("Send Message:\n", text_to_send[:n])
	}
}
