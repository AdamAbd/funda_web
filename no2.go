package main

import "fmt"

func main() {
	var nilaiA, nilaiB float32

	nilaiA = 11.5
	nilaiB = 8.5

	hasil := nilaiA + nilaiB

	fmt.Printf("%.2f + %.2f = %.2f", nilaiA, nilaiB, hasil)
}
