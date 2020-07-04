package main

import "fmt"

func main() {
	q := []int{2, 3, 5, 7, 11, 1231}
	fmt.Println(q)

	r := []bool{true, false, true, true, false, true}
	fmt.Println(r)

	s := []struct {
		i int
		b bool
	}{
		{12, true},
		{34, false},
		{23, true},
	}
	fmt.Println(s)
}
