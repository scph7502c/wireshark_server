package main

import (
	"fmt"
	"net"
)

func main() {
	conn, err := net.Dial(
		"tcp",
		"example.com:80",
	)
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	// warstwa aplikacji
	req := "GET / HTTP/1.1\r\n" +
		"Host: example.com\r\n" +
		"Connection: close\r\n\r\n"

	// enkapsulacja
	_, err = conn.Write([]byte(req))
	if err != nil {
		panic(err)
	}
	// dekapsulacja odbiorcza
	buf := make([]byte, 1024)
	n, err := conn.Read(buf)
	if err != nil {
		panic(err)
	}
	fmt.Println("Otrzymano dekapsulowane: ")
	fmt.Println(string(buf[:n]))

}
