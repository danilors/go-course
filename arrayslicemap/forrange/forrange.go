package main

import "fmt"

func main() {
	numeros := [...]int{1, 2, 3, 4, 5} // defined array with [...]
	fmt.Println(numeros)

	// getting index and array item
	for i, numero := range numeros {
		fmt.Printf("%d) %d\n", i+1, numero)
	}
	// ignoring index value
	for _, numero := range numeros {
		fmt.Printf("numero -> %d\n", numero)
	}

}
