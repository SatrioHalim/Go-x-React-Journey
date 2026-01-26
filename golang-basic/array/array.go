package main

import "fmt"

func main() {

	var angka [3]int = [3]int{10, 20, 30} // max 3 elemen
	fmt.Println("Array : ", angka)
	fmt.Println("Array[1] : ", angka[1])

	/*		
		ARRAY FUNCTION
	*/

	number := [5]int{10, 20, 30, 40, 50}

	// menghitung jumlah elemen array
	fmt.Println("Jumlah elemen array : ", len(number))

	// mengakses elemen array
	fmt.Println("Index ke-1: ", number[1])

	// mengubah nilai elemen array
	number[1] = 80
	fmt.Println("Hasil modifikasi index ke-1: ", number[1])

	// menampilkan seluruh elemen array
	fmt.Println("Showing list of array : ")
	for index,value := range number{
		fmt.Println("Isi index ke-",index," = ", value)
	}

	// Compare Array
	arr1 := [3]int{1,2,3}
	arr2 := [3]int{4,5,6}
	fmt.Println("arr1 == arr2 : ",arr1==arr2)
	fmt.Println("arr1 != arr2 : ",arr1!=arr2)

}