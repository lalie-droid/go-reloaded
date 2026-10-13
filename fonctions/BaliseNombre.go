package Fonctions

import (
	"strconv"
	"strings"
)

// pareil que balise avec nombre
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
