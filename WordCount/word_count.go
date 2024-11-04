package WordCount

import (
	"strings"
)

func WordCount(s string) map[string]int {
	stringToArray := strings.Fields(s)

	m := make(map[string]int)

	// Boucle à travers l'array
	for i := 0; i < len(stringToArray); i++ {
		// On vérifie si la clef est déjà présente
		// Si oui on incrément sa valeur et la réinsigne
		if value, ok := m[stringToArray[i]]; ok {
			value += 1
			m[stringToArray[i]] = value

			continue
		}

		// Sinon on initialise sa valeur à 0
		m[stringToArray[i]] = 1
	}

	return m
}
