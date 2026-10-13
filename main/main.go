package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	fichierEntre := os.Args[1]
	fichierSortie := os.Args[2]

	// Lit le fichier d'entrée et applique les transformations
	// om doit entrer go run . fichier entrer fichier sortie
	content, err := ModifierFichier(fichierEntre)
	if err != nil {
		fmt.Println("Erreur :", err)
		return
	}

	// Écrit le résultat dans le fichier de sortie
	err = EcritFichier(fichierSortie, content)
	if err != nil {
		fmt.Println("Erreur :", err)
		return
	}
}

func ModifierFichier(nomFichier string) (string, error) {
	file, err := os.Open(nomFichier)
	if err != nil {
		return "", fmt.Errorf("erreur d'ouverture du fichier: %v", err)
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		transformLine := transformLine(line)
		lines = append(lines, transformLine)

	}

	if err := scanner.Err(); err != nil {
		return "", fmt.Errorf("erreur de lecture du fichier: %v", err)
	}

	return strings.Join(lines, "\n"), nil
}

// EcritFichier écrit le contenu transformé dans le fichier de sortie.
func EcritFichier(nomFichier string, content string) error {
	file, err := os.Create(nomFichier)
	if err != nil {
		return fmt.Errorf("erreur de création du fichier: %v", err)
	}
	defer file.Close()

	_, err = file.WriteString(content)
	if err != nil {
		return fmt.Errorf("erreur d'écriture dans le fichier: %v", err)
	}

	return nil
}

// transformLine applique toutes les transformations à une ligne de texte.
func transformLine(line string) string {
	line = TransformationBalise(line)
	line = ADevientAn(line)
	line = TransformationBaliseNombre(line)
	line = TransformationPonctuation(line)

	return line
}

// TransformationBalise applique les transformations (hex), (bin), (up), (low), (cap) sur une ligne
func TransformationBalise(line string) string {
	// On separe la ligne en mots
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

				decimal, err := strconv.ParseInt(motPrecedent, 16, 64)
				if err == nil {
					tabMots[i-1] = strconv.Itoa(int(decimal))
					//si pas d'erreurs
					// Supprime (hex) de la liste
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
			if i > 0 {
				motPrecedent := tabMots[i-1]
				tabMots[i-1] = capitalize(motPrecedent)
				tabMots = append(tabMots[:i], tabMots[i+1:]...)
				i--
			}
			//if cap, 2 prend le mot -1,    2fois
		case "(cap,nombre)":
			// Vérifier si le mot commence par "(cap,"
			if strings.HasPrefix(word, "(cap,") {
				nombreStr := strings.TrimPrefix(word, "(cap,")
				nombreStr = strings.TrimSuffix(nombreStr, ")")
				nombre, err := strconv.Atoi(nombreStr)
				if err != nil {
					i++ // Si ce n'est pas un nombre, on passe au mot suivant
					continue
				}

				// Capitaliser les mots précédents
				if i > 0 {
					for a := 0; a < nombre; a++ {
						motPrecedent := tabMots[i-1]
						tabMots[i-1] = capitalize(motPrecedent)
					}
					tabMots = append(tabMots[:i], tabMots[i+1:]...)
					i-- // Recule pour ne pas sauter de mot
				}
			} else {
				i++ // Si ce n'est pas un tag connu, on passe au mot suivant
			}

		}
		i++
	}

	return strings.Join(tabMots, " ")
}

// capitalize convertit un mot avec une majuscule initiale.
func capitalize(word string) string {
	if len(word) == 0 {
		return word
	}
	return strings.ToUpper(string(word[0])) + strings.ToLower(word[1:])
}

// ADevientAn transformation de "a" en "an" et "A" en "An"
func ADevientAn(line string) string {
	//parcours la ligne
	tabMots := strings.Fields(line)
	for i := 0; i < len(tabMots)-1; i++ {
		nextword := tabMots[i+1]
		nextlettre := string(nextword[0])
		if tabMots[i] == "a" {
			//si
			if strings.ContainsAny(nextlettre, "aeiouh") {
				tabMots[i] = "an"
			}
		}
		if tabMots[i] == "A" {

			if strings.ContainsAny(nextlettre, "aeiouh") {
				tabMots[i] = "An"
			}
		}

	}
	return strings.Join(tabMots, " ")
}

func TransformationPonctuation(line string) string {
	tabMots := strings.Fields(line)
	var result string

	for i, word := range tabMots {
		premierlettre := string(word[0])
		// Si le mot est une ponctuation, on ajoute la ponctuation puis un espace
		if word == "..." || word == "?!" {
			result += word
		} else if len(word) == 1 && strings.ContainsAny(word, ".,!?:;") {
			result += word
		} else if len(word) > 1 && strings.ContainsAny(premierlettre, ".,!?:;") {
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

func TransformationBaliseNombre(line string) string {
	tabMots := strings.Fields(line)
	for i := 0; i < len(tabMots); i++ {
		word := tabMots[i]

		switch word {
		case "(up,":
			if i+1 >= len(tabMots) {
				continue
			}
			nbr, err := strconv.Atoi(strings.TrimSuffix(tabMots[i+1], ")"))
			if err != nil || nbr < 0 {
				continue
			}
			debut := i - nbr
			if debut < 0 {
				debut = 0
			}
			for j := debut; j < i; j++ {
				tabMots[j] = strings.ToUpper(tabMots[j])
			}
			tabMots = append(tabMots[:i], tabMots[i+2:]...)
			i--

		case "(low,":
			if i+1 >= len(tabMots) {
				continue
			}
			nbr, err := strconv.Atoi(strings.TrimSuffix(tabMots[i+1], ")"))
			if err != nil || nbr < 0 {
				continue
			}
			debut := i - nbr
			if debut < 0 {
				debut = 0
			}
			for j := debut; j < i; j++ {
				tabMots[j] = strings.ToLower(tabMots[j])
			}
			tabMots = append(tabMots[:i], tabMots[i+2:]...)
			i--

		case "(cap,":
			if i+1 >= len(tabMots) {
				continue
			}
			nbr, err := strconv.Atoi(strings.TrimSuffix(tabMots[i+1], ")"))
			if err != nil || nbr < 0 {
				continue
			}
			debut := i - nbr
			if debut < 0 {
				debut = 0
			}
			for j := debut; j < i; j++ {
				r := []rune(tabMots[j])
				tabMots[j] = strings.ToUpper(string(r[0])) + strings.ToLower(string(r[1:]))
			}
			tabMots = append(tabMots[:i], tabMots[i+2:]...)
			i--
		}
	}
	return strings.Join(tabMots, " ")
}
