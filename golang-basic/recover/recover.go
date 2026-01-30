package main

import "fmt"

func tanganiPanic() {
	if r := recover(); r != nil {
		fmt.Println("Terjadi panic, tapi sudah ditangani : ", r)
	}
}

func bagi(a, b int) {
	defer tanganiPanic() // defer akan memastikan tanganiPanic dipanggil walau terjadi panic
	fmt.Printf("Membagi %d dengan %d\n",a,b)
	hasil := a / b // jika b = 0, akan terjadi panic
	fmt.Println("Hasil: ", hasil)
}

func main() {
	fmt.Println("Program mulai")
	bagi(10,2) // normal
	bagi(10,0) // akan terjadi panic, tapi recover akan menanganinya
	fmt.Println("Program selesai dengan aman")

	/*
		OUTPUT:
	Program mulai
	Membagi 10 dengan 2
	Hasil:  5
	Membagi 10 dengan 0
	Terjadi panic, tapi sudah ditangani :  runtime error: integer divide by zero
	Program selesai dengan aman
	*/
}