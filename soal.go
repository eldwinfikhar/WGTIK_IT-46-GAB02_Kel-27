package main

import "fmt"

type razia struct {
	name, plat, jenis string
	langgar           int
	rute              [10]string
}

var redArea = [6]string{"Gajah Mada", "Hayam Wuruk", "Sisingamangaraja", "Panglima Polim", "Fatmawati", "Tomang Raya"}

func kenaRazia(tanggal int, r razia) int {
	var found, x byte
	var ganjil, genap, tilang bool
	var i, j, langgar int
	langgar = 0
	i = 1
	for {
		found = r.plat[len(r.plat)-i]
		i++
		if found == ' ' {
			break
		}
	}
	x = r.plat[len(r.plat)-i]
	ganjil = x == '1' || x == '2' || x == '5' || x == '7' || x == '9'
	genap = x == '2' || x == '4' || x == '6' || x == '8' || x == '0'
	tilang = (tanggal%2 != 0 && genap) || (tanggal%2 == 0 && ganjil)

	if r.jenis == "Mobil" && tilang {
		for i = 0; i < 6; i++ {
			for j = 0; j < len(r.rute); j++ {
				if r.rute[j] == redArea[i] {
					langgar++
				}
			}
		}
	}

	return langgar
}

func main() {
	var R1, R2, R3, R4 razia

	R1.name = "Denver"
	R1.plat = "B 2791 KDS"
	R1.jenis = "Mobil"
	R1.rute = [10]string{"TB Simatupang", "Panglima Polim", "Depok", "Senen Raya"}
	R1.langgar += kenaRazia(27, R1)

	R2.name = "Toni"
	R2.plat = "B 1212 JBB"
	R2.jenis = "Mobil"
	R2.rute = [10]string{"Pintu Besar Selatan", "Panglima Polim", "Depok", "Senen Raya", "Kemang"}
	R2.langgar += kenaRazia(27, R2)

	R3.name = "Stark"
	R3.plat = "B 444 XSX"
	R3.jenis = "Motor"
	R3.rute = [10]string{"Pondok Indah", "Depok", "Senen Raya", "Kemang"}
	R3.langgar += kenaRazia(27, R3)

	R4.name = "Anna"
	R4.plat = "B 678 DD"
	R4.jenis = "Mobil"
	R4.rute = [10]string{"Fatmawati", "Panglima Polim", "Depok", "Senen Raya", "Kemang", "Gajah Mada"}
	R4.langgar = kenaRazia(27, R4)

	fmt.Print("[")
	if R1.langgar != 0 {
		fmt.Print("{name: ", R1.name, ", tilang: ", R1.langgar, "} ")
	}
	if R2.langgar != 0 {
		fmt.Print("{name: ", R2.name, ", tilang: ", R2.langgar, "}, ")
	}
	if R3.langgar != 0 {
		fmt.Print("{name: ", R3.name, ", tilang: ", R3.langgar, "} ")
	}
	if R4.langgar != 0 {
		fmt.Print("{name: ", R4.name, ", tilang: ", R4.langgar, "}")
	}
	fmt.Print("]")
}
