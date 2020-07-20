package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	file, err := os.Open("./err.go")
	if err != nil {
		log.Fatalln("Error!")
	}
	defer file.Close()
	data := make([]byte, 200)
	count, err := file.Read(data)
	if err != nil {
		log.Fatalln("Error")
	}
	fmt.Println(count, string(data))
}
