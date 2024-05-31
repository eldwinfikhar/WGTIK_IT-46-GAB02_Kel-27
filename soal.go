package main

import "fmt"

func platNo(tanggal int, plat string) bool {
	var found, x byte
	var ganjil, genap, hasil bool
	var i int

	i = 1
	for {
		found = plat[len(plat)-i]
		i++
		if found == '_' {
			break
		}
	}
	x = plat[len(plat)-i]
	ganjil = x == '1' || x == '2' || x == '5' || x == '7' || x == '9'
	genap = x == '2' || x == '4' || x == '6' || x == '8' || x == '0'
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
	var tanggal, i int
	var plat string
	var hasil bool
	// siapkan 3 pasang masukan dengan format:
	// tanggal plat_nomor
	for i = 0; i < 3; i++ {
		fmt.Scanln(&tanggal, &plat)
		hasil = platNo(tanggal, plat)
		if hasil {
			fmt.Println("Tidak Melanggar")
		} else {
			fmt.Println("Melanggar")
		}
	}
}
