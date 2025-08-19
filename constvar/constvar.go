package main

import (
	"fmt"
	"math"
)

func main() {
	const PI float64 = 3.1415
	var raio = 3.2

	//forma reduzida de criar uma var
	area := PI * math.Pow(raio, 2)
	fmt.Println("A área da circunferencia é: ", area)

	const (
		a = 1
		b = 2
	)

	const (
		c = 3
		d = 4
	)

	fmt.Printf("a: %d, b: %d, c: %d, d: %d", a, b, c, d)


}
