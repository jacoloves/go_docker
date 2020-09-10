package main

import "fmt"

func hello() {
	fmt.Println("heello,world")
	return
}

func div(a, b int) (int, int) {
	q := a / b
	r := a % b
	return q, r
}

func funcSmp() {
	hello()
	q, r := div(19, 7)
	fmt.Printf("sho=%d div=%d\n", q, r)
	fmt.Printf("%#v\n", func(x, y, int) int { return x + y })
	fmt.Printf("%#v\n", func(x, y, int) int { return x + y }(2, 3))
}
