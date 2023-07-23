package main

import "fmt"

func isMultiple(cek, factor int) bool {
	return cek%factor == 0
}

func main() {
	var cek int

	var factor int

	fmt.Print("Masukan angka yang mau anda cek : ")
	fmt.Scanln(&cek)

	fmt.Print("Masukan factor kelipatan : ")
	fmt.Scanln(&factor)

	if isMultiple(cek, factor) {
		fmt.Printf("%d adalah kelipatan dari %d\n", cek, factor)
	} else {
		fmt.Printf("%d bukan kelipatan dari %d\n", cek, factor)
	}
}
