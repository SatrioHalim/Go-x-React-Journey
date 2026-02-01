package main

import "fmt"

// Struct
type PersegiPanjang struct {
	panjang float64
	lebar   float64
}

// Method menghitung luas persegi panjang
// ini seperti menambahkan function pada struct
func (p PersegiPanjang) Luas() float64 {
	// (p PersegiPanjang) adalah receiver jadi 
	// bisa mengakses atribut struct
	// tanpa perlu mengirim parameter lagi
	return p.panjang * p.lebar
}

func main(){
	// Membuat objek
	pp := PersegiPanjang{panjang: 10,lebar: 5}
	fmt.Println("Panjang :", pp.panjang)
	fmt.Println("Lebar   :", pp.lebar)
	fmt.Println("Luas    :", pp.Luas())

	/*
		OUTPUT:
		Panjang : 10
		Lebar   : 5
		Luas    : 50
	*/
}