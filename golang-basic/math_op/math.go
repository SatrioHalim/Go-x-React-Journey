package main

import "fmt"

func main() {
	x := 10

	// Assign
	x = 5
	fmt.Println("x = 5 : ",x)
	// Add and Assign
	x += 10
	fmt.Println("x += 10 : ",x)
	// Subtract and Assign
	x -= 5
	fmt.Println("x -= 5 : ",x)
	// Multiply and Assign
	x *= 2
	fmt.Println("x *= 2 : ",x)
	// Divide and Assign
	x /= 3
	fmt.Println("x /= 3 : ",x)
	// Modulus and Assign
	x %= 4
	fmt.Println("x %= 4 : ",x)
	x++ // Increment
	fmt.Println("x++ : ",x)
	x-- // Decrement
	fmt.Println("x-- : ",x)
}