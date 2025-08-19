package main

import "fmt"

func main() {
	i := 0

	for i < 10 {
		fmt.Println("i: ", i)
		i++
	}

	for x := 0; x < 10; x++ {
		fmt.Println("x: ", x)
	}

	// infinity loop
	for {

	}
}
