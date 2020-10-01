package main

import "fmt"

var S = "123"

func init() {
	S = S + "A"
}

func init() {
	S = S + "B"
}

func init() {
	S = S + "C"
}

func intiSmp() {
	fmt.Println(S)
}
