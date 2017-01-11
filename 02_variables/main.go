package main

import "fmt"

func main() {
	a := 42
	b := "golang"
	c := 3.15
	d := true

	var e int
	var f string
	var g bool

	var h, i = "hhh", "iii"
	// var j, k = "jjjkkk" <= does not assign jjjkkk to both, the numbers must match

	fmt.Printf("%v (%T) \n", a, a)
	fmt.Printf("%v (%T) \n", b, b)
	fmt.Printf("%v (%T) \n", c, c)
	fmt.Printf("%v (%T) \n", d, d)

	fmt.Printf("%v (%T) \n", e, e)
	fmt.Printf("%v (%T) \n", f, f)
	fmt.Printf("%v (%T) \n", g, g)

	fmt.Printf("%v (%T) \n", h, h)
	fmt.Printf("%v (%T) \n", i, i)
}
