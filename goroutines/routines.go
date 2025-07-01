package main

import (
	"fmt"
	"time"
)

func print(str string) {
	time.Sleep(1 * time.Second)
	fmt.Println(str)
}

func main() {
	// go routine 1
	go print("hello 0")
	time.Sleep(1 * time.Second)
	// go routine 2
	go print("hello 1")
	time.Sleep(1 * time.Second)

	fmt.Println("hello")

}
