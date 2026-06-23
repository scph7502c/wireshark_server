package main

import (
	"fmt"
	"net"
)

func main() {
	localNetwork := "192.168.200.0/24"
	_, subnet, err := net.ParseCIDR(localNetwork)
	if err != nil {
		panic(err)
	}
	targetIP := net.ParseIP("192.168.200.10")
	if subnet.Contains(targetIP) {
		fmt.Println("Ruch lokalny LAN")
	} else {
		fmt.Println("Ruch kierowany w WAN")
	}

	hostPort := net.JoinHostPort(targetIP.String(), "8080")
	tcpAddr, err := net.ResolveTCPAddr("tcp", hostPort)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Gotowy adres trasowania: %s\n", tcpAddr.String())

}
