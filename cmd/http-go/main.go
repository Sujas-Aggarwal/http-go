package main

import (
	"fmt"

	"github.com/Sujas-Aggarwal/http-go/internal/helper"
	"github.com/Sujas-Aggarwal/http-go/internal/tcp"
)

var PORT string = ":8000"

func main() {
	listener := tcp.StartListener(PORT)
	defer helper.CloseListener(listener)
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Error Accepting Connection")
		}
		go tcp.HandleConnection(conn)
	}
}
