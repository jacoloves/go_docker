package main

import "fmt"

type Vertex struct {
	x, y int
}

var (
	v1 = Vertex{3, 2}
	v2 = Vertex{x: 4}
	v3 = Vertex{}
	p  = &Vertex{10, 6}
)

func main() {
	fmt.Println(v1, p, v2, v3)
}
