package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("Hello world")
	fmt.Println("Hello " + "World")
	fmt.Println(string("Hello World"[0]))
	var s string = "Hello World"

	s = strings.Replace(s, "H", "X", 1)
	fmt.Println(s)
	fmt.Println(strings.Contains(s, "world"))

	fmt.Println(`ts
						ts
ts`)

	fmt.Println("\"")
	fmt.Println(`"`)
}
