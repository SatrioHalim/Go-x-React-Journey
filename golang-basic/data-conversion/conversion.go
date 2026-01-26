package main

import (
	"fmt"
	"strconv"
)

func main() {

	// INT to FLOAT conversion
	var a int = 10
	var b float64 = float64(a)
	fmt.Println("Integer value:", a)
	fmt.Println("Converted to float64:", b)

	// INT to STRING conversion
	var score int = 95
	var text string = strconv.Itoa(score)
	fmt.Println("Nilai ujian : ", text)

	// STRING to INT conversion
	var txt string = "123"
	// var txt2 string = "123abc" // contoh yang invalid
	number, err := strconv.Atoi(txt)
	if err != nil {
		fmt.Println("Error message:", err.Error())
	} else {
		fmt.Println("int:", number)
	}

	// BOOLEAN to STRING conversion
	truth := true
	str := strconv.FormatBool(truth)
	fmt.Println("Boolean Ke String:", str)

	// STRING to BOOLEAN conversion
	val , _ := strconv.ParseBool("true") // _ dipake kalo ga make return error
	fmt.Println("String ke Boolean:", val)


}
