package main

import "fmt"

type Vertex struct {
	Lat, Long float64
}

var m map[string]Vertex

func main() {
	m = make(map[string]Vertex)
	m["Bell Labs"] = Vertex{
		111111111111, 3.1414214356,
	}
	fmt.Println(m["Bell Labs"])
}
