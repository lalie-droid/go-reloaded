package Fonctions

import "strings"

func TransformationPonctuation(line string) string {
	tabMots := strings.Fields(line)
	var result string

	for i, word := range tabMots {
		premierlettre := string(word[0])
		// Si le mot est une ponctuation speciale on la fait en premier
		if word == "..." || word == "?!" {
			result += word
		} else if len(word) == 1 && strings.ContainsAny(word, ".,!?:;") {
			//si le mot est de longueur 1 et est une ponctuation
			result += word
		} else if len(word) > 1 && strings.ContainsAny(premierlettre, ".,!?:;") {
			//si la ponctuation est collee a l'avant d'un autre mot
			result += premierlettre
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

//modification car . venait avant ... et n'affichait que un . a chaque fois
