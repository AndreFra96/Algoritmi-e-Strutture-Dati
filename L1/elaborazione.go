package l1

import (
	"errors"
	"fmt"
)

func Elaborazione() {
	var arr = []int{12, 3, 4, 8, 9, 2}
	fmt.Printf("Strano Prodotto: %d\n", stranoProdotto(arr))
	pariDispari(arr)
	result, err := minDispari(arr)
	fmt.Print("Min dispari: ")
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println(result)
	}

}

// Prodotto fra tutti gli elementi dell'array in input maggiori di 7 e multipli di 4
func stranoProdotto(numeri []int) int {
	prodotto := 1

	for i := 0; i < len(numeri); i++ {
		if numeri[i] > 7 && numeri[i]%4 == 0 {
			prodotto *= numeri[i]
		}
	}

	return prodotto
}

func pariDispari(numeri []int) {
	for i := 0; i < len(numeri); i++ {
		if i%2 == 0 {
			fmt.Printf("%d è pari\n", numeri[i])
		} else {
			fmt.Printf("%d è dispari\n", numeri[i])

		}
	}
}

// Restituisce il più piccolo numero dispari della slice
func minDispari(numeri []int) (int, error) {
	if len(numeri) < 1 {
		return -1, errors.New("numeri is empty")
	}
	min := numeri[0]
	for i := 1; i < len(numeri); i++ {
		if numeri[i]%2 != 0 && numeri[i] < min {
			min = numeri[i]
		}
	}
	return min, nil
}
