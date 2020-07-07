package main

import "fmt"

type Vertex struct {
	Lat, Long float64
}

var m = map[string]Vertex{
	"Bell Labs": {40.342, 123.222}
	"Google": {40.342, 123.222}
}

func main() {
	fmt.Println(m)
}