package main

import "fmt"

// Menamai return value "hasil"
func tambah(a int, b int) (hasil int) {
	hasil = a + b
	// cukup tulis return aja
	return
}

// Void function (tanpa return value)
func sapa(nama string) {
	fmt.Println("Halo,", nama)
}

// Fungsi variadic untuk menjumlahkan banyak angka
func jumlahAngka(angka ...int) int { 
	// range meng-iterasi slice angka, range 
	// mengembalikan dua nilai: index dan value
	total := 0
	for _, v:= range angka {
		total += v
	}
	return total
}

func main() {
	fmt.Println(tambah(3,5)) // output: 8
	fmt.Println(jumlahAngka(1,2,3,4,5)) // output: 15
	sapa("Budi") // output: Halo, Budi
}