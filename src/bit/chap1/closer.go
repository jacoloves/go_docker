package main

import "fmt"

func cicleArea(pi float64) func(radius float64) float64 {
	return func(radius float64) float64 {
		return pi * radius * radius
	}
}

func main() {
	c1 := cicleArea(3.14)
	fmt.Println(c1(2))

	c2 := cicleArea(3)
	fmt.Println(c2(2))
}
