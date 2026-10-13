package Fonctions

import (
	"strconv"
	"strings"
)

// TransformationBalise applique les transformations (hex), (bin), (up), (low), (cap) sur une ligne
func TransformationBalise(line string) string {
	// On sépare la ligne en mots
	tabMots := strings.Fields(line)

	for i := 0; i < len(tabMots); i++ {
		word := tabMots[i]

		switch word {
		case "(hex)":
			// verifie qu'il y a un mot
			if i > 0 {
				motPrecedent := tabMots[i-1]
				// ParseInt fait la convertion automatiquement en toute base

				decimal, err := strconv.ParseInt(motPrecedent, 16, 64)
				if err == nil {
					tabMots[i-1] = strconv.Itoa(int(decimal))
					//si pas d'erreurs
					// Supprime l'instruction "(hex)" de la liste
					tabMots = append(tabMots[:i], tabMots[i+1:]...)
					// On recule l'index pour ne pas sauter de mot
					i--
				}
			}
		case "(bin)":
			if i > 0 {
				motPrecedent := tabMots[i-1]
				decimal, err := strconv.ParseInt(motPrecedent, 2, 64)
				if err == nil {
					tabMots[i-1] = strconv.Itoa(int(decimal))
					tabMots = append(tabMots[:i], tabMots[i+1:]...)
					i--
				}
			}
		case "(up)":
			// met le mot  en majuscules
			if i > 0 {
				motPrecedent := tabMots[i-1]
				tabMots[i-1] = strings.ToUpper(motPrecedent)
				tabMots = append(tabMots[:i], tabMots[i+1:]...)
				i--
			}
		case "(low)":
			// met le mot precedent en minuscules
			if i > 0 {
				motPrecedent := tabMots[i-1]
				tabMots[i-1] = strings.ToLower(motPrecedent)
				tabMots = append(tabMots[:i], tabMots[i+1:]...)
				i--
			}
		case "(cap)":
			//pareil en cap
			if i > 0 {
				motPrecedent := tabMots[i-1]
				if len(motPrecedent) == 0 {
					return motPrecedent
				} else {
					return strings.ToUpper(string(motPrecedent[0])) + strings.ToLower(motPrecedent[1:])
				}

			}
			tabMots = append(tabMots[:i], tabMots[i+1:]...)
			i--
		}
		i++
	}

	return strings.Join(tabMots, " ")
}

// capitalize convertit un mot avec une majuscule initiale.
// jai fait une autre fonction lutot que l'integree a transformation balise pour que ce soit plus lisible
func capitalize(word string) string {
	if len(word) == 0 {
		return word
	}
	return strings.ToUpper(string(word[0])) + strings.ToLower(word[1:])
}
