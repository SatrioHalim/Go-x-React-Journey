package main

import "fmt"

func main() {
	
	defer fmt.Println("defer ini tetap jalan sebelum program mati")
	fmt.Println("Sebelum panic")
	panic("ada sesuatu yang fatal")
	// baris di bawah panic tidak akan dieksekusi
	fmt.Println("Setelah panic")

	/*
	OUTPUT:
	Sebelum panic
	defer ini tetap jalan sebelum program mati
	panic: ada sesuatu yang fatal
	*/
}