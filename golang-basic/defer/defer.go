package main

import "fmt"

func main() {

	// konsep defer : LIFO (Last In First Out)
	defer fmt.Println("A : Dieksekusi paling akhir")
	fmt.Println("B : Dieksekusi paling awal")
	defer fmt.Println("C : Dieksekusi kedua paling akhir")
	fmt.Println("D : Dieksekusi kedua")
	/*
	OUTPUT:
	B : Dieksekusi paling awal
	D : Dieksekusi kedua
	C : Dieksekusi kedua paling akhir
	A : Dieksekusi paling akhir
	*/

	// Contoh kedua
	x := 1
	defer fmt.Println("defer-1, x =", x) // argumen x dievaluasi sekarang (1)
	x = 2
	defer fmt.Println("defer-2, x =", x) // argumen x dievaluasi sekarang (2)
	fmt.Println("body, x=",x)
	/*
	OUTPUT:
	body, x= 2
	defer-2, x = 2
	defer-1, x = 1
	*/

}