package main

import (
	"fmt"

	"github.com/yousefzinsazk78/firstgowebapp/page"
)

func main() {

	p1 := &page.Page{Title: "first title", Body: []byte("yousef is here....")}
	p1.Save()
	p2, _ := page.LoadPage("first title")
	fmt.Println(string(p2.Body))

}
