package main

import (
	"fmt"
	"sort"
)

func main() {
	i := []int{5, 3, 2, 8, 7}
	s := []string{"d", "a", "f"}
	p := []struct {
		Name string
		Age  int
	}{
		{"Nancy", 20},
		{"Vera", 40},
		{"Mike", 30},
		{"Bob", 50},
	}
	fmt.Println(i, s, p)
	sort.Ints(i)
	sort.Strings(s)
	fmt.Println(i, s, p)
}
