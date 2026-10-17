package helper

import (
	"fmt"
	"net"
	"strconv"
	"strings"
)

func Parser(rawHeaders string) RequestHeader {
	// we need to iterate over the headers and get the useful information out of it
	headers := strings.Split(rawHeaders, "\n")
	// we got individual headers now
	// now we need to parse it and get the useful information out of it
	// for the first line, considering its gonna be our method route and versioning, it will be handled differently
	firstLine := strings.Split(headers[0], " ")
	method := firstLine[0]
	route := firstLine[1]
	version := firstLine[2]
	host := ""
	useragent := ""
	accept := ""
	contentlength := 0
	contenttype := ""
	// now we can easily just loop over our headers and get what we want for the remaining of them
	for _, i := range headers[1:] {
		// the i is now a line
		// we need to divide the line in 2 by :\n
		if len(i) < 2 {
			continue
		}
		raw := strings.Split(i, ": ")
		// now we got 2 parts in our raw which are useful to use
		switch raw[0] {
		case "Host":
			host = strings.TrimSpace(raw[1])
		case "User-Agent":
			useragent = strings.TrimSpace(raw[1])
		case "Accept":
			accept = strings.TrimSpace(raw[1])
		case "Content-Length":
			x, err := strconv.Atoi(strings.TrimSpace(raw[1]))
			if err != nil {
				fmt.Println("ERROR: Couln't parse Content Length")
			}
			contentlength = int(x)
		case "Content-Type":
			contenttype = strings.TrimSpace(raw[1])
		}
	}
	// we got our headers now, lets send back the response now -
	return RequestHeader{Host: host, Method: method, Version: version, Route: route, UserAgent: useragent, Accept: accept, ContentLength: contentlength, ContentType: contenttype}
}

func BodyParser(conn net.Conn, remBody []byte, neededBytes int) []byte {
	if neededBytes <= int(len(remBody)) {
		return remBody[:neededBytes]
	}
	// if not then we can simply read more buffer and put that into remaining body byte and return that
	buffer := make([]byte, 1024)
	remaining := neededBytes - len(remBody)
	for remaining > 0 {
		n, err := conn.Read(buffer)
		if err != nil {
			fmt.Println("ERROR: Reading Body of the Request -", err)
			return remBody
		}
		req := min(remaining, n)
		remBody = append(remBody, buffer[:req]...)
		remaining -= req
	}
	return remBody
}
