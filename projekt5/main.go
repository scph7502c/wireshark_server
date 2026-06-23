package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	url := "http://example.com"

	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Odczytano %d bajtów ładunku\n", len(body))
}
