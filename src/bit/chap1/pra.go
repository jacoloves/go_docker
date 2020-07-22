package main

import "fmt"

func main() {
	f := 1.11
	fmt.Println(int(f))

	s := []int{1, 2, 5, 6, 2, 3, 1}
	fmt.Println(s[2:4])

	m := make(map[string]int)
	m["Mike"] = 20
	m["Nancy"] = 24
	m["Messi"] = 30
	fmt.Printf("%T %v\n", m, m)
}
