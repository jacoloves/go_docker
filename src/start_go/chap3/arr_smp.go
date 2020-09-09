package main

import "fmt"

func arrSmp() {
	a := [5]int{1, 2, 3, 4, 5}
	fmt.Printf("%v", a)
	a1 := [...]int{1, 2, 3}
	a1[0] = 0
	a1[2] = 0
	fmt.Printf("%v", a1)
}
