package main

import "fmt"

// prima function untuk cek bilangan prima
func isPrime(prima int) bool {
	if prima < 2 {
		return false
	}

	for i := 2; i < prima; i++ {
		if prima%i == 0 {
			return false
		}
	}
	return true
}

// main funtion
func main() {
	var prima int

	fmt.Print("Masukan bilangan yang ingin anda cek: ")
	fmt.Scanln(&prima)

	if isPrime(prima) {
		fmt.Println(prima, "adalah bilangan prima.")
	} else {
		fmt.Println(prima, "bukan bilangan prima.")
	}
}
