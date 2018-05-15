package main

import "fmt"

func main() {
	x := 24
	fmt.Println(x)
	{
		fmt.Println(x)
		y := "closure example"
		fmt.Println(y)
	}
}
