package main

import "fmt"

func main() {
	fmt.Print("Normal print")
	fmt.Print("Same line")

	fmt.Println("new Line")
	fmt.Println("another new Line")

	x := 3.41565

	xs := fmt.Sprint(x)
	fmt.Println("The value of x is " + xs)
	fmt.Println("The value of x is ", x)

	fmt.Printf("The value of x is %.2f", x) // output //The value of x is 3.42

	a := 1
	b := 1.900
	c := false
	d := "opa"
	fmt.Printf("\n %d %.1f %t %s", a, b, c, d) // output  1 1.9 false opa
	fmt.Printf("\n %v %v %v %v", a, b, c, d)   // output ...
}
