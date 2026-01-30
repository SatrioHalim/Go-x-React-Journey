package main

import "fmt"

// contoh fungsi cleanup sederhana
func cleanup(){
	fmt.Println("Cleanup : Menutup resource...")
}

func bacaConfig(namaFile string){
	// defer selalu dipanggil meskipun ada panic
	defer cleanup()

	// recover diletakkan dalam defer untuk menangani panic
	defer func(){ // anonymous function
		if r:= recover(); r != nil {
			fmt.Println("Terjadi error:", r)
		}
	}()
	if namaFile == "" {
		panic("Nama file konfigurasi tidak boleh kosong")
	}
	fmt.Println("Membaca konfigurasi dari file:", namaFile)
	// anggap berhasil membaca
}

func main(){
	fmt.Println("Case 1")
	bacaConfig("")
	fmt.Println("Program tetap berjalan setelah panic")
	fmt.Println("Case 2")
	bacaConfig("config.yaml")

	/*
		OUTPUT:
	Case 1
	Terjadi error: Nama file konfigurasi tidak boleh kosong
	Cleanup : Menutup resource...
	Program tetap berjalan setelah panic
	Case 2
	Membaca konfigurasi dari file: config.yaml
	Cleanup : Menutup resource...
	*/
}