package main

import "fmt"

func main() {
	angka := 10

	switch angka {
	case 1:
		fmt.Println("Angka adalah Satu")
	case 5:
		fmt.Println("Angka adalah Lima")
	case 10:
		fmt.Println("Angka adalah Sepuluh")
	default:
		fmt.Println("Angka bukan 1, 5, atau 10")
	}

	hari := "Sabtu"

	switch hari {
	case "Senin", "Selasa", "Rabu", "Kamis", "Jumat":
		fmt.Println("Hari Kerja")
	case "Sabtu", "Minggu":
		fmt.Println("Hari Libur")
	default:
		fmt.Println("Bukan hari yang valid")
	}
}