package main

import (
	"fmt"
)

func main() {
	hex := "1a3f"
	decimal := hexToDecimal(hex)
	fmt.Printf("Hexadécimal %s = Décimal %d\n", hex, decimal)
}
