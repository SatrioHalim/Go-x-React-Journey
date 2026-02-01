package main

import "fmt"

func ubahNama(nama *string) {
	*nama = "Golang" // ubah value di alamat memory tersebut
}

func heal(hp *int){
	*hp += 20 // tambah hp sebesar 20 langsung di memory address
	fmt.Println("HP meningkat 20!")
}

func attack(hp *int, damage int){
	*hp -= damage // kurangi hp sebesar damage langsung di memory address
	fmt.Printf("Pemain diserang! HP berkurang %d!\n",damage)
	if (*hp < 0){
		fmt.Println("Game over! Pemain kalah 😵")
	}
}

func main() {
	fmt.Println("Case 1:")
	nama := "Pemrograman"
	fmt.Println("Sebelum diubah:", nama)
	ubahNama(&nama)
	fmt.Println("Sesudah diubah:", nama)
	fmt.Println()
	fmt.Println("Case 2:")
	hp := 50
	fmt.Println("HP awal pemain:", hp)
	
	heal(&hp) // kirim alamat memory hp
	fmt.Println("HP sekarang:", hp)

	attack(&hp,10)
	fmt.Println("HP sekarang:", hp)

	attack(&hp,80)
	fmt.Println("HP sekarang:", hp)
}