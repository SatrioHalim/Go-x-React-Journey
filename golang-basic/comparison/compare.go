package main

import "fmt"

func main() {
	var angka1, angka2 int

	// Scan input user
	fmt.Print("Masukkan angka pertama: ")
	fmt.Scanln(&angka1)
	fmt.Print("Masukkan angka kedua: ")
	fmt.Scanln(&angka2)

	// Compare (true jika kondisi terpenuhi, false jika tidak)
	fmt.Println("\n==== Hasil Perbandingan====")
	fmt.Printf("%d == %d ? %v\n", angka1,angka2, angka1 == angka2)
	fmt.Printf("%d != %d ? %v\n", angka1,angka2, angka1 != angka2)
	fmt.Printf("%d > %d ? %v\n", angka1,angka2, angka1 > angka2)
	fmt.Printf("%d < %d ? %v\n", angka1,angka2, angka1 < angka2)
	fmt.Printf("%d >= %d ? %v\n", angka1,angka2, angka1 >= angka2)
	fmt.Printf("%d <= %d ? %v\n", angka1,angka2, angka1 <= angka2)
}