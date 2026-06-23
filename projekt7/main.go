package main

import (
	"fmt"
	"net"
)

func main() {
	domain := "ciszaart.pl"
	fmt.Printf("Analiza DNS: %s\n\n", domain)
	ips, err := net.LookupIP(domain)
	if err != nil {
		fmt.Println("Błąd LookupIp: ", err)

	} else {
		fmt.Println("---- Rekordy A/AAAA ----")
		for _, ip := range ips {
			fmt.Printf("Znaleziono IP: %s\n", ip.String())
			fmt.Println(ip.MarshalText())
		}
	}
	fmt.Println("\n ---- Rekordy MX ----")
	mxs, err := net.LookupMX(domain)
	if err != nil {
		fmt.Println("Bład LookupMX", err)
	} else {
		for _, mx := range mxs {
			fmt.Printf(
				"Host: %s (Priorytet: %d)\n",
				mx.Host,
				mx.Pref,
			)
		}
	}
	fmt.Println("\n---- Rekordy TXT ----")
	txts, err := net.LookupTXT(domain)
	if err != nil {
		fmt.Println("Błąd LookupTXT", err)
	} else {
		for _, txt := range txts {
			if len(txt) > 40 {
				txt = txt[:37] + "..."
			}
			fmt.Printf("Wartość: %s\n", txt)
		}
	}

}
