package main

import "fmt"

func init() {
	fmt.Println("Init!")
}

func bazz() {
	fmt.Println("fizz")
}

func main() {
	bazz()
	fmt.Println("Hello World", "test test")
}
