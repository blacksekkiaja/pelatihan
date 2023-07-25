package main

import (
	"fmt"
	"strconv"
	"strings"
)

func ArrayMerge(arrayA, arrayB []string) []string {
	merge := append(arrayA, arrayB...)
	nameMap := make(map[string]bool)
	uniqueSlice := []string{}
	for _, name := range merge {
		if !nameMap[name] {
			nameMap[name] = true
			uniqueSlice = append(uniqueSlice, name)
		}
	}
	return uniqueSlice
}

func Mapping(slice []string) map[string]int {
	hitungMuncul := make(map[string]int)

	for _, str := range slice {
		hitung := strings.Count(strings.Join(slice, " "), str)

		hitungMuncul[str] = hitung
	}

	return hitungMuncul
}

func munculSekali(angka string) []int {
	hitungAngka := make(map[int]int)
	for _, num := range angka {
		n, err := strconv.Atoi(string(num))
		if err == nil {
			hitungAngka[n]++
		}
	}

	var angkaUnik []int
	for a, hitung := range hitungAngka {
		if hitung == 1 {
			angkaUnik = append(angkaUnik, a)
		}
	}

	return angkaUnik
}

func main() {
	// Test case soal 1
	fmt.Println("hasil dari soal 1 'merge array'")

	fmt.Println(ArrayMerge([]string{"king", "devil jin", "akuma"}, []string{"eddie", "steve", "geese"}))

	fmt.Println(ArrayMerge([]string{"sergei", "jin"}, []string{"jin", "steve", "bryan"}))

	fmt.Println(ArrayMerge([]string{"alisa", "yoshimitsu"}, []string{"devil jin", "yoshimitsu", "alisa", "law"}))

	fmt.Println(ArrayMerge([]string{}, []string{"devil jin", "sergey"}))

	fmt.Println(ArrayMerge([]string{"hwoarang"}, []string{}))

	fmt.Println(ArrayMerge([]string{}, []string{}))

	// Test case soal 2
	fmt.Println()
	fmt.Println("hasil dari soal 2 'hitung jumlah kata yang keluar'")

	fmt.Println(Mapping([]string{"asd", "qwe", "asd", "adi", "qwe", "qwe"}))

	fmt.Println(Mapping([]string{}))

	// test case soal 3
	fmt.Println()
	fmt.Println("hasil dari soal 3 'angka yang muncul satu kali'")

	fmt.Println(munculSekali("1234123"))
	fmt.Println(munculSekali("76523752"))
	fmt.Println(munculSekali("12345"))
	fmt.Println(munculSekali("1122334455"))
	fmt.Println(munculSekali("0872504"))

}
