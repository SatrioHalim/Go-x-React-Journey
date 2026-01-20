package main

import "fmt"

func main(){
	// Deklarasi dengan kutip ganda
	nama := "Halim"
	pesan := "Selamat datang tapi boong!"

	// Deklarasi dengan backtick (raw string literal) -> ditampil apa adanya termasuk enter dan spasi
	paragraf := `lorem ipsum dolor sit amet, consectetur adipiscing elit.
	sed do eiusmod tempor incididunt ut labore et dolore magna aliqua.`

	// Output
	fmt.Println("Nama :", nama)
	fmt.Println("Pesan :", pesan)
	fmt.Println("Paragraf :\n", paragraf)
}