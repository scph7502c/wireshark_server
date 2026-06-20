package main

import (
	"fmt"
	"net"
	"os"
)

func handleNode(c net.Conn) {
	defer c.Close()
	remoteAddr := c.RemoteAddr()
	fmt.Println("Połączono z: ", remoteAddr)
	msg := []byte("Witaj w Gwieździe!\n")
	_, err := c.Write(msg)
	if err != nil {
		fmt.Println("Bład zapisu: ", err)
		return
	}
}

// star topology
func main() {
	addr := "127.0.0.1:8181"
	listener, err := net.Listen(
		"tcp",
		addr,
	)
	if err != nil {
		fmt.Println("Błąd bindowania:", err)
		os.Exit(1)
	}
	defer listener.Close()
	fmt.Println("Węzeł centralny działa.")
	fmt.Println("Oczekuję na węzły sieci...")
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Błąd: ", err)
			continue
		}
		go handleNode(conn)

	}

}
