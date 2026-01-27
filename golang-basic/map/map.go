package main

import "fmt"

func main() {
	mahasiswa := map[string]string{
		"nama":  "Halim",
		"kelas": "LA01",
	}
	fmt.Println("Nama : ",mahasiswa["nama"])
	fmt.Println("Kelas : ",mahasiswa["kelas"])

	// menambahkan key
	mahasiswa["jurusan"] = "Informatika"
	fmt.Println("Jurusan : ",mahasiswa["jurusan"])

	mahasiswa2 := map[string]int{
		"Halim": 90,
		"Mr Ngarap": 100,
		"Julius": 80,
	}
	// menghapus key
	delete(mahasiswa2, "Julius")

	// mengakses semua key dan value pada map
	for key, value := range mahasiswa2 {
		fmt.Println(key, " : ", value)
	}
}