package main

import "fmt"

func main() {
	l := []int{100, 300, 23, 11, 2, 4, 6, 4}

	min := 100000100000
	for _, num := range l {
		if min > num {
			min = num
		}
	}

	fmt.Println(min)
}
