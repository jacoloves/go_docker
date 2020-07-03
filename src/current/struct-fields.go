package main

import "fmt"

type Vertex struct {
	X int
	Y int
}

func main() {
	v := Vertex{3, 9}
	v.X = 8
	v.Y = 1223
	fmt.Println(v.X)
	fmt.Println(v.Y)
}
