package fonctions

import "fmt"

func hexToDecimal(hex string) int {
	decimal := 0
	for _, char := range hex {
		switch {
		case char >= '0' && char <= '9':
			decimal = decimal*16 + int(char-'0')
		case char >= 'a' && char <= 'f':
			decimal = decimal*16 + int(char-'a') + 10
		case char >= 'A' && char <= 'F':
			decimal = decimal*16 + int(char-'A') + 10
		default:
			fmt.Printf("Caractère invalide dans l'hexadécimal: %c\n", char)
			return -1 // Retourne -1 en cas d'erreur
		}
	}
	return decimal
}
