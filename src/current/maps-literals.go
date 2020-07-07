package main

import "fmt"

type Vertex struct {
	Lat, Long float64
}

var m = map[string]Vertex{
	"Bell Labs":Vertex{
		40.68433, -74.38867,
	},
	"Google":Vertex{
		100.342, 111.234,
	},
	"Benson":Vertex{
		1.1111, 2.222,
	},
}


func main() {
	fmt.Println(m)
	fmt.Println(m["Bell Labs"])
	fmt.Println(m["Google"])
	fmt.Println(m["Benson"])
}
