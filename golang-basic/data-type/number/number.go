package main

import(
	"fmt"
)

func main(){

	// signed integer
	var i8 int8 = 127
	var i16 int16 = 32767
	var i32 int32 = 2147483647
	var i64 int64 = 9223372036854775807
	var i int = -100 // int is either 32 or 64 bit depending on the system architecture

	fmt.Println("Signed Integer Types:")
	// Kalau mau pakai Printf bisa pakai verb %v atau placeholder lainnya
	fmt.Printf("int8 : %v\n", i8)
	fmt.Printf("int16 : %v\n", i16)
	fmt.Printf("int32 : %v\n", i32)
	fmt.Printf("int64 : %v\n", i64)
	fmt.Printf("int : %v\n", i)

	// unsigned integer
	var u8 uint8 = 255
	var u16 uint16 = 65535
	var u32 uint32 = 4294967295
	var u64 uint64 = 18446744073709551615
	var u uint = 100 // uint is either 32 or 64 bit depending on the system architecture

	fmt.Println("\nUnsigned Integer Types:")
	fmt.Println("uint8 :", u8)
	fmt.Println("uint16 :", u16)
	fmt.Println("uint32 :", u32)
	fmt.Println("uint64 :", u64)
	fmt.Println("uint :", u)

	// floating-point numbers
	var f32 float32 = 3.1415
	var f64 float64 = 3.141592653589793

	fmt.Println("\nFloating-Point Types:")
	fmt.Println("float32 :", f32)
	fmt.Println("float64 :", f64)
}