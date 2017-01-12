package main

import "fmt"

func zero(y *int) {
	*y = 0
}

func main() {
	x := 42
	fmt.Println("value of x before:\t", x)

	zero(&x)
	fmt.Println("value of x after:\t", x)
}
