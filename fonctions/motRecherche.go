package fonctions

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func MotRecherche(MotATrouver string) (bool, error) {
	fichier, err := os.Open("sample.txt")
	if err != nil {
		return false, fmt.Errorf("Erreur d'ouverture: %v", err)
	}
	defer fichier.Close()
	scanner := bufio.NewScanner(fichier)
	for scanner.Scan() {
		ligne := scanner.Text()
		mots := strings.Fields(ligne)
		for _, mot := range mots {
			if strings.EqualFold(mot, MotATrouver) {
				return true, nil
			}
		}
	}

	if err := scanner.Err(); err != nil {
		return false, fmt.Errorf("erreur lors de la lecture du fichier : %v", err)
	}

	return false, nil

}
