package main

import "fmt"

func main() {
	i, j := 42, 100

	p := &i
	fmt.Println(*p)
	*p = 3
	fmt.Println(i)

	p = &j
	*p = *p / 4
	fmt.Println(j)
}
