package fonctions

import (
	"fmt"
	"os"
)

func copie() {
	contenu, err := os.ReadFile("sample.txt")
	if err != nil {
		fmt.Println("Erreur de lecture :", err)
		return
	}

	err = os.WriteFile("copie.txt", contenu, 0644)
	if err != nil {
		fmt.Println("Erreur d'écriture :", err)
		return
	}

}
