package main

import "fmt"

func hitungLuasTrapesium(a, b, h float64) float64 {
	luas := 0.5 * (a + b) * h
	return luas
}

func main() {
	var sisiAtas, sisiBawah, tinggi float64

	fmt.Print("Masukkan panjang sisi atas trapesium: ")
	fmt.Scanln(&sisiAtas)

	fmt.Print("Masukkan panjang sisi bawah trapesium: ")
	fmt.Scanln(&sisiBawah)

	fmt.Print("Masukkan tinggi trapesium: ")
	fmt.Scanln(&tinggi)

	luasTrapesium := hitungLuasTrapesium(sisiAtas, sisiBawah, tinggi)
	fmt.Printf("Luas trapesium adalah: %.2f\n", luasTrapesium)
}
