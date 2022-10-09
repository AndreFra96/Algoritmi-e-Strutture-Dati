package l1

import "fmt"

func Conto_corrente() {

	//Raccolgo saldo iniziale
	var saldo int
	fmt.Println("Saldo iniziale:")
	fmt.Scanf("%d", &saldo)

	//Raccolgo spese
	var spesa int
	for {
		fmt.Println("Inserire spesa:")
		fmt.Scanf("%d", &spesa)
		//Sottraggo spesa al saldo
		saldo -= spesa
		//Se il saldo è <= 0 interrompo
		if saldo <= 0 {
			break
		}
	}

	//Stampo saldo finale
	fmt.Printf("Saldo finale: %d €\n", saldo)

}
