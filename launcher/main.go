package main

import (
	"fmt"
	"l1"
)

func main() {
	fmt.Printf("Cosa vuoi eseguire: \n1) Supera 100\n2) Conto Corrente\n3) Andamento\n4) Elaborazione\n5) Elaborazione Stringhe\n6) Sottosequenze Crescenti\n7) Ottieni sottostringhe\n8) Permutazione da ordinare\n9) Pesce Lanterna\n")
	var scelta int
	fmt.Scanf("%d", &scelta)

	switch scelta {
	case 1:
		l1.Supera100()
	case 2:
		l1.Conto_corrente()
	case 3:
		l1.Andamento()
	case 4:
		l1.Elaborazione()
	case 5:
		l1.ElaborazioneStringhe()
	case 6:
		l1.SottosequenzeCrescenti()
	case 7:
		l1.SottostringheAB()
	case 8:
		l1.PermutazioneDaOrdinare()
	case 9:
		l1.PesceLanterna()
	default:
		l1.Hello_world()
	}
}
