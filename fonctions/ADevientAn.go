package Fonctions

import "strings"

// ADevientAn transformation de "a" en "an" et "A" en "An"
func ADevientAn(line string) string {
	//parcours la ligne
	tabMots := strings.Fields(line)
	for i := 0; i < len(tabMots)-1; i++ {
		//fait une variable pour la premier lettre du prochain mot
		nextword := tabMots[i+1]
		nextlettre := string(nextword[0])
		if tabMots[i] == "a" {
			//si le mot est a on verifie que la prochaine lettre est aeiouh
			if strings.ContainsAny(nextlettre, "aeiouh") {
				tabMots[i] = "an"
			}
		}
		//pareil en maj
		if tabMots[i] == "A" {

			if strings.ContainsAny(nextlettre, "aeiouh") {
				tabMots[i] = "An"
			}
		}

	}
	return strings.Join(tabMots, " ")
}

// correction appliquée car a devenait an a chaque mot comprenant une voyelle
