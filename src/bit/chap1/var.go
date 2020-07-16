package main

import "fmt"

var (
	i    int     = 2
	f64  float64 = 2.3
	s    string  = "test"
	t, f bool    = true, false
)

func foo() {
	xi := 1
	var xf32 float32 = 23.3
	xs := "ttt"
	xt, xf := true, false
	fmt.Println(xi, xf32, xs, xt, xf)
	fmt.Printf("%T\n", xf32)
	fmt.Printf("%T\n", xi)
}

func main() {
	fmt.Println(i, f64, s, t, f)
	foo()
}
