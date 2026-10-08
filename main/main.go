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
	// Étape 1 : Gérer les apostrophes
	line = TransformationBalise(line)

	// Étape 4 : Gérer "a" -> "an"
	line = ADevientAn(line)

	return line
}

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
			// Met le mot  en majuscules
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
	tabMots := strings.Fields(line)
	for i := 0; i < len(tabMots)-1; i++ {
		if tabMots[i] == "a" {
			nextMot := tabMots[i+1]
			if strings.ContainsAny(nextMot, "aeiouh") {
				tabMots[i] = "an"
			}
		}
		if tabMots[i] == "A" {
			nextMot := tabMots[i+1]
			if strings.ContainsAny(nextMot, "aeiouh") {
				tabMots[i] = "An"
			}
		}

	}
	return strings.Join(tabMots, " ")
}
