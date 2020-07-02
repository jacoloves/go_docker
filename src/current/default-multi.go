package main

import "fmt"

func main() {
	fmt.Println("test")

	for i := 0; i < 11; i++ {
		defer fmt.Println(i)
	}

	fmt.Println("test2")
}
