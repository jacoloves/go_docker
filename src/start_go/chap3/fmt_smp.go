package main

import (
	"fmt"
)

func fmtsmp() {
	fmt.Println("---------------start---------------")
	fmt.Println("Hello, Golang")
	fmt.Println("My", "name", "is", "Taro")
	fmt.Printf("num=%d\n", 5)
	fmt.Printf("dec=%d binary=%b 8shin=%o 16shin=%x\n", 17, 17, 17, 17)
	fmt.Printf("%d/%d/%d/\n", 2015, 12)
	fmt.Printf("%d/%d/%d\n", 2015, 12, 17)

	fmt.Printf("num=%v str=%v arr=%v\n", 5, "golang", [...]int{1, 2, 3})
	fmt.Printf("num=%#v str=%#v arr=%#v\n", 5, "golang", [...]int{1, 2, 3})
	fmt.Printf("num=%T str=%T arr=%T\n", 5, "golang", [...]int{1, 2, 3})

	print("Hello, WOrld!")
	println("Hello, WOrld!")
	fmt.Println("---------------end---------------")
}
