package main

import (
	"bufio"
	"fmt"
	"net"
)

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:8181")
	if err != nil {
		fmt.Println("NIe udało się połączyć: ", err)
		return
	}
	defer conn.Close()

	message, err := bufio.NewReader(conn).ReadString('\n')
	if err != nil {
		fmt.Println("Błąd odczytu: ", err)
	}
	fmt.Println("Otrzymano z serwera: " + message)
}
