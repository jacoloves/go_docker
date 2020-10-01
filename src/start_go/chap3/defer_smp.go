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

func panictest() {
	defer fmt.Println("on defer")

	panic("runtine error!")
	fmt.Println("hello, world!")
}

func recovertest() {
	defer func() {
		if x := recover(); x != nil {
			fmt.Println(x)
		}
	}()
	panic("panic!")
	fmt.Println("hello, world!")
}

func testRecover(src interface{}) {
	defer func() {
		if x := recover(); x != nil {
			switch v := x.(type) {
			case int:
				fmt.Printf("panic: int=%v\n", v)
			case string:
				fmt.Printf("panic: string=%v\n", v)
			default:
				fmt.Println("panic: unknown")
			}
		}
	}()
	panic(src)
	return
}

func deferSmp() {
	testRecover(128)
	testRecover("hogehoge")
	testRecover([...]int{1, 2, 3})
}
