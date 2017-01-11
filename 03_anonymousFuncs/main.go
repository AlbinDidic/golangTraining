package main

import "fmt"

func main() {
	x := 42

	// Func expression (asignin a func {anounymous} to a variable)
	increment := func() int {
		x++
		return x
	}

	fmt.Println(increment())
	fmt.Println(increment())
	fmt.Println(increment())
}
