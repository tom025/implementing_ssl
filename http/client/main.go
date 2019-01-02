package main

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"net"
)

const httpPort = 80

func main() {
	host := "google.com"
	path := "/"

	addrs, err := net.LookupIP(host)
	if err != nil {
		log.Fatal(err)
	}

	addr, err := firstIPv4(addrs)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("*   Trying %+v...\n", addr)

	conn, err := net.DialTCP(
		"tcp",
		nil,
		&net.TCPAddr{IP: addr, Port: httpPort},
	)
	defer func() {
		err = conn.Close()
		if err != nil {
			log.Fatalf("Error closing connection: %s", err.Error())
		}
	}()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("* Connected to %s (%s) port %d\n", host, addr, httpPort)

	fmt.Fprintf(conn, "GET %s HTTP/1.1 \r\n", path)
	fmt.Printf("> GET %s HTTP/1.1 \r\n", path)

	fmt.Fprintf(conn, "Host: %s\r\n", host)
	fmt.Printf("> Host: %s\r\n", host)

	fmt.Fprintf(conn, "Connection: close\r\n")
	fmt.Printf("> Connection: close\r\n")

	fmt.Fprintf(conn, "\r\n")
	fmt.Printf("> \r\n")

	scanner := bufio.NewScanner(conn)
	var uptoBody bool
	for scanner.Scan() {
		text := scanner.Text()
		if !uptoBody {
			fmt.Printf("< %s\n", text)
		} else {
			fmt.Printf("%s\n", text)
		}
		uptoBody = text == "" || uptoBody
	}
	if scanner.Err() != nil {
		log.Fatal(err)
	}
}

func firstIPv4(addrs []net.IP) (net.IP, error) {
	for _, addr := range addrs {
		ip := addr.To4()
		if ip != nil {
			return ip, nil
		}
	}
	return nil, errors.New("No IPv4 addresses found")
}
