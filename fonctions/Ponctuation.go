package Fonctions

import "strings"

func TransformationPonctuation(line string) string {
	tabMots := strings.Fields(line)
	var result string

	for i, word := range tabMots {
		// Si le mot est une ponctuation, on ajoute la ponctuation puis un espace
		if len(word) == 1 && strings.ContainsAny(word, ".,!?:;") {
			result += word
		} else if word == "..." || word == "?!" {
			result += word
		} else {
			// Sinon, on ajoute un espace avant le mot (sauf pour le premier mot)
			if i > 0 {
				result += " "
			}
			result += word
		}
	}

	return result
}
