package l1

import "fmt"

func Supera100() {
	var i int
	var supera100 int
	fmt.Scanf("%d", &i)
	for i != -1 {
		if supera100 == 0 && i > 100 {
			supera100 = i
		}
		fmt.Scanf("%d", &i)
	}
	if supera100 == 0 {
		fmt.Println("nessun numero maggiore di 100")
	} else {
		fmt.Println(supera100)
	}

}
