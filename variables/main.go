package main

import "fmt"

func main() {

	var a int
	var b string
	var c float64
	var d bool

	fmt.Printf("Variable Declaration assigned to zero value \n")
	fmt.Printf("%v \n", a)
	fmt.Printf("%v \n", b)
	fmt.Printf("%v \n", c)
	fmt.Printf("%v \n", d)

	w := 400
	x := "hello"
	y := 50.0
	z := true

	fmt.Printf("\nVariable assign and initialize\n")
	fmt.Printf("%v \n", w)
	fmt.Printf("%v \n", x)
	fmt.Printf("%v \n", y)
	fmt.Printf("%v \n", z)
}
