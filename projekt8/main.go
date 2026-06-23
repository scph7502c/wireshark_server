package main

import (
	"fmt"
	"io"
	"net"
	"time"
)

func main() {
	listener, err := net.Listen(
		"tcp",
		"127.0.0.1:0",
	)
	if err != nil {
		panic(err)
	}
	defer listener.Close()
	fmt.Println("Nasłuch na: ", listener.Addr())
	done := make(chan struct{})

	//serwer (accept)
	go func() {
		defer close(done)
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Błąd Accept: ", err)
			return
		}
		defer conn.Close()

		buf := make([]byte, 1024)
		n, err := conn.Read(buf)
		if err != nil {
			if err != io.EOF {
				fmt.Println("Błąd odczytu: ", err)
			}
			return
		}
		fmt.Printf("Serwer odebrał: %s\n", buf[:n])

	}()
	// klient (Dial)
	conn, err := net.Dial(
		"tcp",
		listener.Addr().String(),
	)
	if err != nil {
		panic(err)
	}
	msg := []byte("Dane przez TCP")
	_, err = conn.Write(msg)
	if err != nil {
		panic(err)
	}
	time.Sleep(100 * time.Millisecond)
	conn.Close()
	<-done
	fmt.Println("Sesja TCP zakończona.")

}
