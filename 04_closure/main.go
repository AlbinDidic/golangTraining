package main

import "fmt"

func wrapper() func() int {
	x := 42

	return func() int {
		x++
		return x
	}
}

func wrapperParam() func(int) int {
	x := 42

	return func(y int) int {
		x++
		return x + y
	}
}

func main() {
	increment := wrapper()
	incrementParam := wrapperParam()

	fmt.Println(increment())       // 43 <= x got incremented
	fmt.Println(incrementParam(2)) // 45 <= x got incremented to 43 then added y which is 2
	fmt.Println(incrementParam(2)) // 46 <= x was now 43 from previous operation, got incremented to 44 then added y which is 2
}
