package main

import (
	"bufio"
	"fmt"
	"go-reloaded/Fonctions"
	"os"
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
	line = Fonctions.TransformationBalise(line)
	line = Fonctions.ADevientAn(line)
	line = Fonctions.TransformationBaliseNombre(line)
	//line = TransformationGuillemets(line)
	line = Fonctions.TransformationPonctuation(line)

	return line
}
