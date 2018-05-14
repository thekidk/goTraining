package main

import "fmt"

func max(x int) int {
	return 42 + x
}

func main() {
	num := max(7)
	fmt.Println(num)
}