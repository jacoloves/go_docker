package main

import "fmt"

func arrSmp() {
	a := [5]int{1, 2, 3, 4, 5}
	fmt.Printf("%v", a)
	a1 := [...]int{1, 2, 3}
	a1[0] = 0
	a1[2] = 0
	fmt.Printf("%v", a1)

	b1 := [3]int{1, 2, 3}
	b2 := [3]int{4, 5, 6}

	b1 = b2

	fmt.Printf("%v\n", b1)

	b1[0] = 0
	b1[2] = 0

	fmt.Printf("%v\n", b1)
	fmt.Printf("%v\n", b2)

	var x interface{}
	fmt.Printf("%#v", x)
}
