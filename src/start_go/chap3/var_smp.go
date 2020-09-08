package main

import (
	"fmt"
	"math"
)

var n = 100

func varsmp() {
	fmt.Println("---------------start---------------")
	n = n + 1
	fmt.Printf("n=%d\n", n)

	ui_1 := uint32(400000000)
	ui_2 := uint32(4000000000)
	sum := ui_1 + ui_2
	fmt.Println("%d + %d = %d\n", ui_1, ui_2, sum)
	fmt.Printf("uint32 max value = %d\n", math.MaxUint32)
	fmt.Printf("value = %v\n", 1.000000000000000)
	fmt.Printf("value = %v\n", 1.000000000000001)
	fmt.Printf("value = %v\n", 1.000000000000002)
	fmt.Printf("value = %v\n", 1.000000000000003)
	fmt.Printf("value = %v\n", 1.000000000000004)
	fmt.Printf("value = %v\n", 1.000000000000005)
	fmt.Printf("value = %v\n", 1.000000000000006)
	fmt.Printf("value = %v\n", 1.000000000000007)
	fmt.Printf("value = %v\n", 1.000000000000008)
	fmt.Printf("value = %v\n", 1.000000000000009)
	fmt.Println("---------------end---------------")
}
