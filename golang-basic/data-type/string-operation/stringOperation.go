package main

import (
	"fmt"
	"strings"
)

func main(){
	// Concatenation
	// str := "Hello, " + "Golang!"
	// fmt.Println(str)

	// String length
	// length := "Golang"
	// fmt.Println("Length of string:", len(length))

	// Accessing characters
	// char := "Golang"
	// fmt.Println("First character:", string(char[0])) 
	// fmt.Println("Second character:", string(char[1]))

	// Convert to Lowercase
	str := "GOLANG"
	fmt.Println("Lowercase:", strings.ToLower(str))

	// Convert to Uppercase
	str2 := "golang"
	fmt.Println("Uppercase:", strings.ToUpper(str2))

	// prefix
	str3 := "Golang is great"
	fmt.Println("Starts with Golang? : ", strings.HasPrefix(str3,"Golang"))

	// Contains
	str4 := "I love Golang programming"
	fmt.Println("Contains 'Golang'? : ", strings.Contains(str4,"Golang"))

	// Split
	str5 := "apple,banana,cherry"
	fruits := strings.Split(str5, ",")
	fmt.Println("Fruits:", fruits)

	// Replace
	str6 := "I love Java"
	newStr := strings.ReplaceAll(str6, "Java", "Golang")
	fmt.Println("After replacement:", newStr)
}