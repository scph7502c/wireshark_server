package main

import (
	"fmt"
	"io"
	"net"
)

func main() {
	conn, err := net.Dial(
		"tcp",
		"192.168.50.6:6379",
	)
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	// warstwa aplikacji
	req := "GET / HTTP/1.1\r\n" +
		"Host: 192.168.50.6\r\n" +
		"Connection: close\r\n\r\n"

	// enkapsulacja
	_, err = conn.Write([]byte(req))
	if err != nil {
		panic(err)
	}
	// dekapsulacja odbiorcza
	buf := make([]byte, 1024)
	fmt.Println("Otrzymano dekapsulowane: ")

	for {
		n, err := conn.Read(buf)
		if n > 0 {
			fmt.Print(string(buf[:n]))
		}
		if err != nil {
			if err == io.EOF {
				break
			}
			panic(err)
		}

	}

}
