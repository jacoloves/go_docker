package main

import "fmt"

func lenSmp() {
	s := make([]int, 8)
	fmt.Println(len(s))
	a := [3]int{1, 2, 3}
	fmt.Println(len(a))
}

func capSmp() {
	s1 := make([]int, 5)
	fmt.Println(len(s1))
	fmt.Println(cap(s1))
	s2 := make([]int, 5, 10)
	fmt.Println(len(s2))
	fmt.Println(cap(s2))
}

func slaSmp() {
	s := []int{1, 2, 3, 4, 5}
	fmt.Println(len(s))
	fmt.Println(cap(s))
}

func easySla() {
	a := [5]int{1, 2, 3, 4, 5}
	s := a[0:3]
	fmt.Println(s)

	a1 := [...]int{1, 2, 3, 4, 5}
	fmt.Println(a1[len(a1)-2:])

	// s1 := []string{"A", "B", "C"}
	// fmt.Println(s1[0:5])

	s1 := "ABCDE"[1:3]
	fmt.Println(s1)
}

func sliceSmp() {
	// s := make([]int, 10)
	// fmt.Println(s)
	// var a [10]int
	// fmt.Println(a)
	// a := make([]float64, 3)
	// fmt.Println(a)
	// a[0] = 3.14
	// fmt.Println(a)
	// a[1] = 6.28
	// fmt.Println(a)
	// fmt.Println(a[0])
	// fmt.Println(a[4])
	lenSmp()
	capSmp()
	slaSmp()
	easySla()
}
