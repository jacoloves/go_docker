package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	_ "log"
)

func main() {
	r := bytes.NewBuffer([]byte("abc"))
	content, _ := ioutil.ReadAll(r)
	fmt.Println(string(content))
}
