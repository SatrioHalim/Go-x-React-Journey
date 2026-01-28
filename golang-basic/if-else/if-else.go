package main

import "fmt"

func main() {
	var nilai int
	fmt.Print("Masukkan nilai ujian : ")
	fmt.Scanln(&nilai)

	if(nilai >= 90){
		fmt.Println("Nilai anda A")
	} else if(nilai >= 75){
		fmt.Println("Nilai anda B")
	} else if(nilai >= 60){
		fmt.Println("Nilai anda C")
	} else {
		fmt.Println("Nilai anda perlu perbaikan")
	}
}