package main

import (
	"fmt"
)

func main() {
	fmt.Println("here")

	a := addMulti(1, 1, 1, 1, 1, 1)

	fmt.Println(a)

	// anonymous function declared to a variable
	a1 := func(a, b int) (c int) {
		c = a + b
		return
	}

	fmt.Println(a1(1, 2))

	// anonymous immediate function
	func(name string) {
		fmt.Println("aaa ", name)
	}("a")

	// anonymous function go routines
	go func(name string) {
		fmt.Println("my name is ", name)
	}("balaji")

	// closure function has another function as an output
	adding := add(5)
	adding2 := adding(2)
	adding3 := adding2(3)
	fmt.Println("here", adding3)
	fmt.Println()

}

// arbitrary functions
// can have n number of arguments for same dataType by mentioning ...
func addMulti(a ...int) (total int) {

	for _, v := range a {
		fmt.Println("v ", v)
		total += v
	}
	return
}

// closure function
func add(c int) func(int) func(int) int {

	a := func(y int) func(int) int {
		return func(b int) int {
			return c + y + b
		}
	}
	return a
}
