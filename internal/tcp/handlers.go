package tcp

import (
	"bytes"
	"fmt"
	"log"
	"net"

	"github.com/Sujas-Aggarwal/http-go/internal/helper"
)

func StartListener(port string) net.Listener {
	listener, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalln("ERROR: Error starting TCP Listener -", err)
	}
	fmt.Println("INFO: Successfully Started Listening on", port)
	return listener
}

func HandleConnection(conn net.Conn) {
	defer helper.CloseConnection(conn)
	fmt.Println("INFO: Got New Connection -", conn)

	buffer := make([]byte, 1024)
	var accumulated []byte // Track raw bytes to avoid string conversion loss

	for {
		n, err := conn.Read(buffer)
		if err != nil {
			fmt.Println("ERROR: Reading Request Data -", err)
			return
		}

		// Append the newly read chunk
		accumulated = append(accumulated, buffer[:n]...)

		// Check if the delimiter exists anywhere in what we've read so far
		if idx := bytes.Index(accumulated, []byte("\r\n\r\n")); idx != -1 {
			// 1. Split headers and the trailing data (start of the body)
			headerBytes := accumulated[:idx+4]
			trailingBodyBytes := accumulated[idx+4:]

			// 2. Convert headers to string and call your parser
			headersStr := string(headerBytes)
			requestHeader := helper.Parser(headersStr)
			remainingBodyLength := requestHeader.ContentLength

			if remainingBodyLength > 0 {
				// Calculate how much more body data we still need to read
				neededBytes := remainingBodyLength

				// Combine what we already have with a fresh read loop for the rest
				fullBody := helper.BodyParser(conn, trailingBodyBytes, neededBytes)

				// Process your body here
				fmt.Println("Headers:\n", requestHeader)
				fmt.Println("Body", string(fullBody))
				// now we need to process the Request and send the response back.
			}

			// Break the loop since the request ingestion is done
			// we need to expose the final Request to the user so that he can do some processing on it and actually use it later for his own use case
			// and send some response back
			break
		}
	}
}
