package l1

import (
	"fmt"
	"strings"
	"time"
)

func SottostringheAB() {
	word := "cdbaagbabbgbabkasjbfjkfhabskjfbbajsbdsbdbsfbajksfblaskbljbwqbajslblasjfb"
	start := time.Now()
	ottieniSottostringhe(word, "a", "b")
	elapsed := time.Since(start)
	fmt.Printf("Doppia iterazione: %s\n", elapsed)
	start = time.Now()
	count := ottieniSottostringheOpt(word, "a", "b")
	elapsed = time.Since(start)
	fmt.Printf("Singola iterazione: %s\n", elapsed)
	fmt.Println(count)
}

// Doppia iterazione
func ottieniSottostringhe(word string, start string, end string) int {
	splittedWord := strings.Split(word, "")
	var wordsCount int
	for index, char := range splittedWord {
		//Se sono all'inizio di una stringa
		if char == start {
			//Scorro tutte le lettere dopo quella
			for i := index + 1; i < len(splittedWord); i++ {
				//Se trovo una finale aggiungo 1 al count
				if splittedWord[i] == end {
					wordsCount++
				}
			}
		}
	}
	return wordsCount
}

// Singola iterazione
func ottieniSottostringheOpt(word string, start string, end string) int {
	splittedWord := strings.Split(word, "")
	var startCount int
	var wordsCount int
	for _, char := range splittedWord {
		//Se ho una stringa di inizio
		if char == start {
			// aggiungo 1 alle stringhe iniziate
			startCount++
			//Se ho una stringa di fine
		} else if char == end {
			// aggiungo il numero di parole iniziate al numero di parole totali
			wordsCount += startCount
		}
	}
	return wordsCount
}
