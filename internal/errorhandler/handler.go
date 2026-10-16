package errorhandler

import (
	"errors"
	"fmt"
	"io"
	"log"
	"net"
)

func HandleIoError(err error) bool {
	if err != nil {
		// 1. Check for a genuine, clean close
		if errors.Is(err, io.EOF) {
			fmt.Println("Client disconnected cleanly (EOF). Closing connection loop.")
			return true // Exit the loop safely without crashing the server
		}

		// 2. Check for an abrupt client disconnect (e.g., app crashed, network cut)
		// On Linux/macOS this is usually "connection reset by peer"
		var opErr *net.OpError
		if errors.As(err, &opErr) {
			log.Printf("Network/OS error occurred: %v\n", opErr.Err)
			return true
		}

		// 3. Handle any other generic reading errors
		log.Printf("Unexpected error reading connection: %v\n", err)
		return true
	}
	return false
}
