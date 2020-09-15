package main

import (
	"fmt"
	"math"
)

const ONE = 1

func one() (int, int) {
	const TWO = 2
	return ONE, TWO
}

const (
	ap1 = "goor"
	ap2 = "toot"
	ap3 = "ert"
)

func morn(m string) {
	fmt.Println(m)
}

func constSmp() {
	x, y := one()
	fmt.Println("x=%d, y=%d\n", x, y)
	f32 := float32(math.Pi)
	f64 := float64(math.Pi)
	fmt.Printf("%v", f32)
	fmt.Printf("%v", f64)
	morn(ap1)
}
