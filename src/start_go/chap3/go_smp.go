package main

import (
	"fmt"
	"runtime"
)

func sub() {
	for {
		fmt.Println("sub loop")
	}
}

func goSmp() {
	// go sub()
	// for {
	// 	fmt.Println("main loop")
	// }
	go fmt.Println("Yeah!")
	fmt.Printf("NumCPU: %d\n", runtime.NumCPU())
	fmt.Printf("NumGoroutine: %d\n", runtime.NumGoroutine())
	fmt.Printf("Version: %s\n", runtime.Version())
}
