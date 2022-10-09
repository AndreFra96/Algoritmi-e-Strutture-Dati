package l1

import "fmt"

func PesceLanterna() {
	//pesci[i] con 0 <= i <= 7 è il numero di pesci attualmente
	//presenti nello stadio i

	pesci := parsePesci(data())
	fmt.Println(pesci)
	for i := 0; i < 256; i++ {
		passaGiorno(&pesci)
	}
	fmt.Println(pesci)
	fmt.Println(contaPesci(&pesci))

}

func parsePesci(data []int) [9]int {
	arr := [9]int{}
	for i := 0; i < len(data); i++ {
		arr[data[i]]++
	}
	return arr
}

func contaPesci(pesci *[9]int) int {
	sum := 0
	for i := 0; i < 9; i++ {
		sum += pesci[i]
	}
	return sum
}

func passaGiorno(pesci *[9]int) {
	//Sposto lo stadio di tutti i pesci un giorno indietro
	next := pesci[6]
	temp := 0
	for i := 5; i >= 0; i-- {
		temp = pesci[i]
		pesci[i] = next
		next = temp
	}
	pesci[6] = temp

	//Sposto indietro i pesci negli stati 7 e 8
	pesci[6] += pesci[7]
	pesci[7] = pesci[8]

	//Aggiungo i nuovi pesci
	pesci[8] = temp

}

func data() []int {
	return []int{5, 4, 3, 5, 1, 1, 2, 1, 2, 1, 3, 2, 3, 4, 5, 1, 2, 4, 3, 2, 5, 1, 4, 2, 1, 1, 2, 5, 4, 4, 4, 1, 5, 4, 5, 2, 1, 2, 5, 5, 4, 1, 3, 1, 4, 2, 4, 2, 5, 1, 3, 5, 3, 2, 3, 1, 1, 4, 5, 2, 4, 3, 1, 5, 5, 1, 3, 1, 3, 2, 2, 4, 1, 3, 4, 3, 3, 4, 1, 3, 4, 3, 4, 5, 2, 1, 1, 1, 4, 5, 5, 1, 1, 3, 2, 4, 1, 2, 2, 2, 4, 1, 2, 5, 5, 1, 4, 5, 2, 4, 2, 1, 5, 4, 1, 3, 4, 1, 2, 3, 1, 5, 1, 3, 4, 5, 4, 1, 4, 3, 3, 3, 5, 5, 1, 1, 5, 1, 5, 5, 1, 5, 2, 1, 5, 1, 2, 3, 5, 5, 1, 3, 3, 1, 5, 3, 4, 3, 4, 3, 2, 5, 2, 1, 2, 5, 1, 1, 1, 1, 5, 1, 1, 4, 3, 3, 5, 1, 1, 1, 4, 4, 1, 3, 3, 5, 5, 4, 3, 2, 1, 2, 2, 3, 4, 1, 5, 4, 3, 1, 1, 5, 1, 4, 2, 3, 2, 2, 3, 4, 1, 3, 4, 1, 4, 3, 4, 3, 1, 3, 3, 1, 1, 4, 1, 1, 1, 4, 5, 3, 1, 1, 2, 5, 2, 5, 1, 5, 3, 3, 1, 3, 5, 5, 1, 5, 4, 3, 1, 5, 1, 1, 5, 5, 1, 1, 2, 5, 5, 5, 1, 1, 3, 2, 2, 3, 4, 5, 5, 2, 5, 4, 2, 1, 5, 1, 4, 4, 5, 4, 4, 1, 2, 1, 1, 2, 3, 5, 5, 1, 3, 1, 4, 2, 3, 3, 1, 4, 1, 1}
}
