package l1

import "fmt"

func Andamento() {
	var attuale int
	var precedente int

	for{
		fmt.Scanf("%d",&attuale)
		if attuale == 0{break}
		if precedente != 0{
			if attuale >= precedente{
				fmt.Println("+")
			}else{
				fmt.Println("-")
			}
		}
		precedente = attuale
	}
}
