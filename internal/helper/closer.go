package helper

import (
	"fmt"
	"net"
)

func CloseListener(listener net.Listener) {
	fmt.Println("INFO: Closing TCP Listener.")
	listener.Close()
}

func CloseConnection(conn net.Conn) {
	fmt.Println("INFO: Closing Connection.")
	conn.Close()
}
