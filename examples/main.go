package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"os"

	"github.com/lioneagle/echarts/charts/option"
)

type A struct {
	X string
}

type Book struct {
	Name   string  `json:"name"`
	Price  float32 `json:"price"`
	Ok     bool    `json:"ok"`
	Desc   string  `json:"desc,omitempty"`
	Author *Author `json:"author,omitempty"`
}
type Author struct {
	Gender int `json:"gender,omitempty"`
	Age    int `json:"age,omitempty"`
}

func main() {
	a := &A{X: "test"}

	tpl, err := template.New("x1").Parse("my___x__{{.X}}__x__")
	if err != nil {
		fmt.Printf("err = %s\n", err)
		return
	}

	err = tpl.Execute(os.Stdout, a)
	if err != nil {
		fmt.Printf("err = %s\n", err)
		return
	}

	var book Book
	book.Name = "testBook"
	book.Ok = true
	bookByte, _ := json.Marshal(book)
	fmt.Printf("%s\n", string(bookByte))

	titleOpts := &option.TitleOpts{Title: "abc"}
	buf, _ := json.Marshal(titleOpts)
	fmt.Printf("%s\n", string(buf))
}
