package main

import "fmt"

func main() {
	a := 42

	fmt.Println("value of a:\t", a)
	fmt.Println("address of a:\t", &a)

	fmt.Println()

	b := &a // var b *int = &a

	fmt.Println("value of b:\t", b)
	fmt.Println("address of b:\t", &b)

	fmt.Println()

	fmt.Println("value on address where b is pointing to: ", *b) // Dereferencing of b

	*b = 69 // Derefrence b and assign a new value. In other words, assign a value on the address that b is pointing to

	fmt.Println()
	fmt.Println("value of a:\t", a)

	// Everything in go is 'pass by value'
}
