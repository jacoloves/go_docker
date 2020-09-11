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

func plus(x, y int) int {
	return x + y
}

var plusAlias = plus

func returnFunc() func() {
	return func() {
		fmt.Println("I'm a function")
	}
}

func callFunction(f func()) {
	f()
}

func later() func(string) string {
	var store string

	return func(next string) string {
		s := store
		store = next
		return s
	}
}

func integers() func() int {
	i := 0
	return func() int {
		i += 1
		return i
	}
}

func funcSmp() {
	hello()
	q, r := div(19, 7)
	fmt.Printf("sho=%d div=%d\n", q, r)
	plusAlias(10, 5)
	f := returnFunc()
	f()
	callFunction(func() {
		fmt.Println("I'm func2")
	})

	f2 := later()

	fmt.Println(f2("golang"))
	fmt.Println(f2("is"))
	fmt.Println(f2("awesome"))

	ints := integers()

	fmt.Println(ints())
	fmt.Println(ints())
	fmt.Println(ints())

	otherInts := integers()
	fmt.Println(otherInts())
}
