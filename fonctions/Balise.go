package fonctions

import (
	"strconv"
	"strings"
)

// TransformationBalise applique les transformations (hex), (bin), (up), (low), (cap) sur une ligne
func TransformationBalise(line string) string {
	// On sépare la ligne en mots
	tabMots := strings.Fields(line)
	i := 0
	for i < len(tabMots) {
		word := tabMots[i]

		switch word {
		case "(hex)":
			// verifie qu'il y a un mot
			if i > 0 {
				motPrecedent := tabMots[i-1]
				// ParseInt fait la convertion automatiquement en toute base
				//pourquoi on met decimal, err ?

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
			// Même logique que pour (hex), mais pour le binaire
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
			// Met le mot précédent en majuscules
			if i > 0 {
				motPrecedent := tabMots[i-1]
				tabMots[i-1] = strings.ToUpper(motPrecedent)
				tabMots = append(tabMots[:i], tabMots[i+1:]...)
				i--
			}
		case "(low)":
			// Met le mot précédent en minuscules
			if i > 0 {
				motPrecedent := tabMots[i-1]
				tabMots[i-1] = strings.ToLower(motPrecedent)
				tabMots = append(tabMots[:i], tabMots[i+1:]...)
				i--
			}
		case "(cap)":
			// Met une majuscule initiale au mot précédent
			if i > 0 {
				motPrecedent := tabMots[i-1]
				tabMots[i-1] = capitalize(motPrecedent)
				tabMots = append(tabMots[:i], tabMots[i+1:]...)
				i--
			}
		}
		i++
	}

	// Reconstruit la ligne à partir des mots transformés
	// a voir
	return strings.Join(tabMots, " ")
}

// capitalize convertit un mot avec une majuscule initiale.
func capitalize(word string) string {
	if len(word) == 0 {
		return word
	}
	return strings.ToUpper(string(word[0])) + strings.ToLower(word[1:])
}
