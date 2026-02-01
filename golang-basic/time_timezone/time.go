package main

import (
	"fmt"
	"time"
)

// Fungsi untuk konversi waktu ke zona tertentu
func convertToZone(t time.Time, zone string) time.Time{
	loc, err := time.LoadLocation(zone)
	if err != nil {
		fmt.Println("Error memuat lokasi:",err)
		return t
	}
	return t.In(loc)
}

func main(){
	// 1. Ambil waktu saat ini dalam UTC
	nowUTC := time.Now().UTC()
	fmt.Println("Waktu saat ini (UTC):", nowUTC)

	// 2. Konversi ke zona waktu lokal (Asia/Jakarta)
	jakartaTime := convertToZone(nowUTC, "Asia/Jakarta")
	fmt.Println("Waktu di Jakarta:", jakartaTime)

	// 3. Konversi ke zona waktu lain yang menggunakan DST misal New York (America/New_York)
	// DST adalah Daylight Saving Time
	newYorkTime := convertToZone(nowUTC, "America/New_York")
	fmt.Println("Waktu di New York:", newYorkTime)

	// 4. Menyimpan waktu (simulasi penyimpanan ke DB)
	// Disarankan simpan dalam UTC
	fmt.Println("Menyimpan waktu ke database (UTC):", nowUTC)

	/*
		OUTPUT:
		Waktu saat ini (UTC): 2024-06-15 08:30:00 +0000 UTC
		Waktu di Jakarta: 2024-06-15 15:30:00 +0700 WIB
		Waktu di New York: 2024-06-15 04:30:00 -0400 EDT
		Menyimpan waktu ke database (UTC): 2024-06-15 08:30:00 +0000 UTC
	*/

}