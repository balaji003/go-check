package main

import "fmt"

func main() {
	fmt.Println("here")

	a := addMulti(1, 1, 1, 1, 1, 1)

	fmt.Println(a)

}

func addMulti(a ...int) (total int) {

	for _, v := range a {
		fmt.Println("v ", v)
		total += v
	}
	return
}
