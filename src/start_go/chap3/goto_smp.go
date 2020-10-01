package main

import "fmt"

func gotoSmp() {
	fmt.Println("A")
	goto L
L:
	fmt.Println("C")
}
