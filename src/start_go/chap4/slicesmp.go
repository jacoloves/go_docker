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

	s2 := "あいうえお"[3:9]
	fmt.Println(s2)

	s3 := []int{1, 2, 3}
	s3 = append(s3, 4)
	fmt.Println(s3)

	s3 = append(s3, 5, 6, 7)
	fmt.Println(s3)

	s4 := []int{1, 2, 3}
	s5 := []int{4, 5, 6}

	s6 := append(s4, s5...)
	fmt.Println(s6)

	var b []byte
	b = append(b, "あいうえお"...)
	b = append(b, "かきくけこ"...)
	b = append(b, "さしすせそ"...)
	fmt.Println(b)

	s7 := make([]int, 0, 0)
	fmt.Printf("(A) len=%d, cap=%d\n", len(s7), cap(s7))

	s7 = append(s7, 1)
	fmt.Printf("(B) len=%d, cap=%d\n", len(s7), cap(s7))

	s7 = append(s7, []int{2, 3, 4}...)
	fmt.Printf("(C) len=%d, cap=%d\n", len(s7), cap(s7))

	s7 = append(s7, 5)
	fmt.Printf("(D) len=%d, cap=%d\n", len(s7), cap(s7))

	s7 = append(s7, 6, 7, 8, 9)
	fmt.Printf("(E) len=%d, cap=%d\n", len(s7), cap(s7))

	s8 := make([]int, 1024, 1024)
	fmt.Printf("len=%d, cap~%d\n", len(s8), cap(s8))

	s8 = append(s8, 0)
	fmt.Printf("len=%d, cap~%d\n", len(s8), cap(s8))
}

func copySmp() {
	s1 := []int{1, 2, 3, 4, 5}
	s2 := []int{10, 11}
	n := copy(s1, s2)
	fmt.Println(n)
	fmt.Println(s1)

	s3 := []int{1, 2, 3, 4, 5}
	s4 := []int{10, 11, 12, 13, 14, 15, 16}
	n2 := copy(s3, s4)

	fmt.Println(n2)
	fmt.Println(s3)

	b := make([]byte, 9)
	n3 := copy(b, "あいうえお")
	fmt.Println(n3, b)

	a := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	s5 := a[2:4]
	fmt.Println(len(s5))
	fmt.Println(cap(s5))

	s6 := a[2:4:4]
	fmt.Println(len(s6))
	fmt.Println(cap(s6))

	s7 := a[2:4:6]
	fmt.Println(len(s7))
	fmt.Println(cap(s7))

	s8 := []string{"Apple", "Banana", "Cherry"}

	for i, v := range s8 {
		fmt.Printf("[%d] => %s\n", i, v)
	}

	for i2 := 0; i2 < len(s8); i2++ {
		fmt.Printf("[%d] -> %s\n", i2, s8[i2])
		s8 = append(s8, "Melon")
	}

	for i, v := range s8 {
		fmt.Printf("{")
	}
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
	copySmp()
}
