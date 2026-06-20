package main

import (
	"fmt"
	"io"
	"net"
	"time"
)

func startServer() {
	addr := "127.0.0.1:8282"
	l, err := net.Listen(
		"tcp",
		addr,
	)
	if err != nil {
		panic(err)
	}
	defer l.Close()
	for {
		c, err := l.Accept()
		if err != nil {
			continue
		}
		go handleLatency(c)
	}
}

func handleLatency(c net.Conn) {
	defer c.Close()
	buf := make([]byte, 1024)
	n, err := c.Read(buf)
	if err != nil {
		if err != io.EOF {
			fmt.Println("Błąd odczytu: ", err)
		}
		return
	}
	fmt.Printf("Otrzymano %d bajtów\n", n)
	time.Sleep(50 * time.Millisecond)
	resp := []byte("Zrozumiano")
	_, err = c.Write(resp)
	if err != nil {
		fmt.Println("Błąd zapisu: ", err)
		return
	}

}

func startClient() {
	start := time.Now()
	c, err := net.Dial(
		"tcp",
		"127.0.0.1:8282",
	)
	if err != nil {
		panic(err)
	}
	defer c.Close()

	_, _ = c.Write([]byte("Dane"))
	buf := make([]byte, 1024)
	_, _ = c.Read(buf)
	latency := time.Since(start)
	fmt.Printf("Całkowite opóźnienie: %v\n", latency)

}

func main() {
	go startServer()
	time.Sleep(10 * time.Millisecond)
	startClient()
}
