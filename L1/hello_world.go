package l1

import "fmt"

func Hello_world() {

	x := "questa è una stringa"
	fmt.Println(x)

	var y string = "seconda stringa"
	fmt.Println(y)

	//Stringa non inizializzata (default '')
	var z string
	fmt.Println(z == "")

	//Dichiarazione array (lunghezza calcolata automaticamente)
	var arr = [...]int{1, 2, 3}
	arr2 := [...]int{4, 5, 6}
	fmt.Println(arr)
	fmt.Println(arr2)

	//Dichiarazione array (lunghezza inserita)
	var arr3 = [3]int{}
	var arr4 = [3]int{7, 8}
	fmt.Println(arr3)
	fmt.Println(arr4)
	fmt.Println(len(arr4)) //Capacità totale

	//Slices
	var slice = []int{}
	var slice2 = append(slice, 1) //La modifica non avviene in-place
	slice3 := []int{}
	slice4 := []int{9, 10, 11}
	slice4 = append(slice4, 1)
	sliceFromArray := arr4[0:1] //Slice creato con gli elementi da 0 a 1 dell'arr4

	fmt.Println(slice)
	fmt.Println(slice2)
	fmt.Println(slice3)
	fmt.Println(slice4)
	fmt.Println(len(slice4)) //Numero di elementi contenuti
	fmt.Println(cap(slice4)) //Capacità totale (ogni volta che la slice è piena e faccio una push viene raddoppiata capacità)
	fmt.Println(sliceFromArray)

	//Slice with make
	sliceWithMake := make([]int, 4, 100) //slice di lunghezza 4 e capacità 100

	fmt.Println(sliceWithMake)
	fmt.Println(len(sliceWithMake))
	fmt.Println(cap(sliceWithMake))

	//Operatori binari
	a := 9
	b := 8
	fmt.Printf("a = %b\n", a)
	fmt.Printf("b = %b\n", b)
	fmt.Printf("a & b is %b\n", a&b)   //AND
	fmt.Printf("a | b is %b\n", a|b)   //OR
	fmt.Printf("a ^ b is %b\n", a|b)   //XOR
	fmt.Printf("a << 2 is %b\n", a<<2) //SLL
	fmt.Printf("a >> 2 is %b\n", a>>2) //SLR

	//Loops
	for i := 0; i <= 5; i++ {
		fmt.Println(i)
	}

	//Range operator (iterazione array,slices e mappe)
	array := [5]string{"a", "b", "c", "d", "e"}
	for indice, value := range array {
		fmt.Printf("Indice: %d, Valore: %s\n", indice, value)
	}
	for _, value := range array {
		fmt.Printf("Valore: %s\n", value)
	}

}
