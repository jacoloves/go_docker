package main

import (
	"io/ioutil"
	"fmt"
)

type Page struct {
	Title string
	Body []byte
}

func (p *Page) save() errer {
	filename := p.Title + ".txt"
	return ioutil.WriteFile(filename, p.Body, 0600)
}

func loadPage(title string) (*Page, error){
	filename := title + ".txt"
	body, err := ioutil.ReadFIle(filename)
	if err != nil {
		return nil, err
	}
	return
