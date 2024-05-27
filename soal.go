package main

import "fmt"

func platNo(tanggal int, plat string) bool {
	var x string
	var ganjil, genap, hasil bool
	x = plat[4:5]
	ganjil = x == "1" || x == "3" || x == "5" || x == "7" || x == "9"
	genap = x == "2" || x == "4" || x == "6" || x == "8" || x == "0"
	if tanggal%2 != 0 && ganjil {
		hasil = true
	} else if tanggal%2 == 0 && genap {
		hasil = true
	} else {
		hasil = false
	}
	return hasil
}

func main() {
	fmt.Print("Hello world")
}
