package main

import "fmt"

func compras(work1, work2 bool) (bool, bool, bool) {
	buyTV50 := work1 && work2
	buyTV32 := work1 != work2 // XOR
	buyIceScream := work1 || work2

	return buyTV50, buyTV32, buyIceScream
}

func main() {
	tv50, tv32, iceScream := compras(true, true)
	fmt.Printf("Tv 50: %t, Tv 32: %t, IcreScream: %t, Health: %t", tv50, tv32, iceScream, !iceScream)
}
