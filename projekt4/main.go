package main

import (
	"fmt"
)

// correctAddressing demonstruje
// prawidlowy dostep do pamieci
// poprzez indeksy wycinka.
func correctAddressing() {
	// 1. Alokacja tablicy bazowej.
	var array = [13]int{
		10, 20, 30, 40, 50,
	}

	// 2. Utworzenie wycinka.
	// Wskazuje na array[12] i array[14].
	sliceFromArr := array[1:3]

	// 3. PRAWIDŁOWE ADRESOWANIE.
	// Używamy operatora indeksu .
	// Wpisujemy wartość 99 do
	// pierwszej komórki, którą
	// widzi wycinek (czyli array[12]).

	// Modyfikacja drugiej komórki
	// (czyli array[14]).
	sliceFromArr[1] = 88

	// Konsumujemy wycinek
	fmt.Printf(
		"Wycinek: %v\n",
		sliceFromArr,
	)

	// Konsumujemy tablicę bazową,
	// by udowodnić współdzielenie
	// tej samej pamięci.
	fmt.Printf(
		"Tablica: %v\n",
		array,
	)

	// 4. Utworzenie bufora bajtów.
	buf := make([]byte, 4)

	// Adresowanie komórek bufora.

	buf[3] = 128

	fmt.Printf(
		"Bufor: %v\n",
		buf,
	)
}

func main() {
	correctAddressing()
}
