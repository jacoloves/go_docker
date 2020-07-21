package main

import (
	"fmt"
	"time"
)

func getOsName() string {
	return "rrrrr"
}

func main() {
	switch os := getOsName(); os {
	case "po":
		fmt.Println("ppp")
	case "win":
		fmt.Println("sss")
	default:
		fmt.Println("ddddd")
	}

	t := time.Now()
	fmt.Println(t.Hour())
	switch {
	case t.Hour() < 12:
		fmt.Println("Mornig")
	case t.Hour() < 17:
		fmt.Println("Afternoon")
	default:
		fmt.Println("night")
	}
}
