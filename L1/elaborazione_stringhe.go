package l1

import (
	"fmt"
	"strings"
)

func ElaborazioneStringhe() {
	arr := []string{"le", "parole", "con", "a", "come", "abbello"}
	fmt.Println(arr)
	fmt.Printf("Quante con a: %d\n", quanteConA((arr)))
	fmt.Printf("Prima con a: %s\n", primaConA((arr)))
	fmt.Printf("Più corta: %d\n", piuCorta((arr)))

}

func quanteConA(parole []string) int {
	var count int
	for i := 0; i < len(parole); i++ {
		if strings.Contains(strings.ToUpper(parole[i]), "A") {
			count++
		}
	}
	return count
}

func primaConA(parole []string) string {
	var parola string
	for _, el := range parole {
		if strings.Contains(strings.ToUpper(el), "A") {
			parola = el
			break
		}
	}
	return parola
}

func piuCorta(parole []string) int {
	if len(parole) < 1 {
		return 0
	}
	minSize := len(parole[0])

	for i := 1; i < len(parole); i++ {
		if len(parole[i]) < minSize {
			minSize = len(parole[i])
		}
	}

	return minSize
}
