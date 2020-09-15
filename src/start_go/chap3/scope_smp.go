package main

import (
	"fmt"
)

func anything(a interface{}) {
	fmt.Println(a)
}

func scopeSmp() {
	fmt.Println("hello, world")
	x, y := 3, 5
	if n := x * y; n%2 == 0 {
		fmt.Printf("n(%d) is even\n", n)
	} else {
		fmt.Printf("n(%d) is odd\n", n)
	}

	for i := 0; i < 100; i++ {
		if i%2 == 1 {
			continue
		}
		fmt.Println(i)
		i++
	}

	for i, r := range "ABC" {
		fmt.Printf("[%d] -> %d\n", i, r)
	}

	for i, r := range "あいうえお" {
		fmt.Printf("[%d] -> %d\n", i, r)
	}

	n := 3
	switch n {
	case 1, 2:
		fmt.Println("1 or 2")
	case 3, 4:
		fmt.Println("3 or 4")
	default:
		fmt.Println("unknown")
	}

	s := "A"
	switch s {
	case "A":
		s += "B"
		fallthrough
	case "B":
		s += "C"
		fallthrough
	case "C":
		s += "D"
		fallthrough
	default:
		s += "E"
	}
	fmt.Println(s)

	n1 := 1
	switch n1 {
	case 1:
		fmt.Println("one")
	case 2.0:
		fmt.Println("two")
	case 3 + 0i:
		fmt.Println("three")
	}

	switch n2 := 2; n2 {
	case 1, 3, 5, 7, 9:
		fmt.Printf("%d is odd\n", n2)
	case 2, 4, 6, 8, 10:
		fmt.Printf("%d is even\n", n2)
	}

	n3 := 4
	switch {
	case n3 > 0 && n3 < 3:
		fmt.Println("0 < n < 3")
	case n3 > 3 && n3 < 6:
		fmt.Println("3 < n < 6")
	}

	anything(1)
	anything(3.14)
	anything(4 + 5i)
	anything('海')
	anything("日本語")
	anything([...]int{1, 2, 3, 4, 5})

}
