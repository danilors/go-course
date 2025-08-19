package main

import "fmt"

func main() {
	i := 1

	var p *int = nil //null value

	p = &i

	*p++
	i++

	fmt.Println("memory pointer ", p) // 0xc000010130
	fmt.Printf("i: %v p: %v", i, *p)  // 3

}
