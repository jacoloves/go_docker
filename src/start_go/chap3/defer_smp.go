package main

import "fmt"

func runDefer() {
	defer fmt.Println("defer")
	fmt.Println("done")
}

func runDefer2() {
	defer fmt.Println("1")
	defer fmt.Println("2")
	defer fmt.Println("3")
	fmt.Println("done")
}

func deferSmp() {
	runDefer2()
}
