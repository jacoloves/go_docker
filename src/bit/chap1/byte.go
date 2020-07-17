package main

import "fmt"

func main() {
	b := []byte{82, 83}
	fmt.Println(b)
	fmt.Println(string(b))

	c := []byte("PR")
	fmt.Println(c)
	fmt.Println(string(c))
}
