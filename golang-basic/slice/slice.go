package main

import "fmt"

func main() {
	arr := [6]int{10, 20, 30, 40, 50, 60}

	s1 := arr[:] // slice all array
	fmt.Println("Ini adalah s1 : ", s1)

	s2 := arr[:3] // slice index 0-2
	fmt.Println("Ini adalah s2 : ", s2)

	s3 := arr[2:] // slice index 2-end
	fmt.Println("Ini adalah s3 : ", s3)

	s4 := arr[1:4] // slice index 1-3
	fmt.Println("Ini adalah s4 : ", s4)

	/*
		SLICE FUNCTION
	*/

	// membuat slice dengan fungsi make
	s := make([]int,3,5) // len=3, capacity=5
	fmt.Println("Slice s : ", s)
	fmt.Println("Len : ", len(s)) // menghitung jumlah elemen slice
	fmt.Println("Capacity : ", cap(s)) // menghitung kapasitas slice

	// menambahkan elemen pada slice
	s = append(s, 10,20)
	fmt.Println("Setelah di append : ", s)

	// copy slice
	slice2 := make([]int,3)
	slice3 := copy(slice2,s) // copy slice s ke slice2
	fmt.Println("copied : ", slice2)
	fmt.Println("Jumlah elemen yang tersalin : ",slice3)

	angka := []int{1,2,3,4,5}
	slice4 := angka[1:4]
	fmt.Println("Slice 4 : ",slice4)
}