package Fonctions

import "strings"

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
