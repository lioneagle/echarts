package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"os"

	"github.com/lioneagle/echarts/charts/option"
	"github.com/lioneagle/echarts/types"
	"github.com/lioneagle/goutil/src/buffer"
)

type A struct {
	X string
}

type Book struct {
	Name   string   `json:"name"`
	Price  float32  `json:"price"`
	Ok     bool     `json:"ok"`
	Desc   string   `json:"desc,omitempty"`
	Author *Author  `json:"author,omitempty"`
	Data   []string `jason:"xx,omitempty"`
}
type Author struct {
	Gender int `json:"gender,omitempty"`
	Age    int `json:"age,omitempty"`
}

type ITest interface {
	A()
}

type Test1 struct {
	x int
}

func (this Test1) A() {
	this.x++
}

type Test2 struct {
	x int
}

func (this *Test2) A() {
	this.x--
}

type Node struct {
	Name  string
	Value string
}

type ExtendOpts struct {
	X    int `json:"id"`
	Data []*Node
}

func (this *ExtendOpts) MarshalJSON() ([]byte, error) {
	//buf1, err := json.Marshal(*this)
	//fmt.Printf("buf1 = %s\n", buf1)
	//return buf1, err

	fmt.Println("MarshalJSON1")
	buf := buffer.NewByteBuffer(nil)
	buf.WriteByte('{')

	for i, v := range this.Data {
		fmt.Fprintf(buf, "\"%s\": {\"x\":%s}", v.Name, v.Value)
		if i < len(this.Data)-1 {
			buf.WriteString(",")
		} else {
			buf.WriteString("\n")
		}
	}
	buf.WriteByte('}')
	fmt.Printf("buf = %s\n", buf.String())

	//panic("xx1")
	return buf.Bytes(), nil
	//return []byte(`{"x1": 100,"x2":101,"x3":102}`), nil
}

//*/

func (this *ExtendOpts) AddNode(nodes ...*Node) {
	this.Data = append(this.Data, nodes...)
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
	fmt.Println("")

	var book Book
	book.Name = "testBook"
	book.Ok = true
	book.Data = make([]string, 0)
	bookByte, _ := json.Marshal(book)
	fmt.Printf("json(bookByte) = %s\n", string(bookByte))

	titleOpts := &option.TitleOpts{Title: "abc"}
	buf, err := json.Marshal(titleOpts)
	if err != nil {
		fmt.Printf("err = %#v\n", err)
	}
	fmt.Printf("%s\n", string(buf))

	var set1 types.StringSet

	fmt.Println("find ret =", set1.Find("a"))

	t1 := Test1{}
	t2 := Test2{}

	var itest1 ITest
	var itest2 ITest

	itest1 = t1
	itest2 = &t2

	t1.A()
	(&t1).A()
	t2.A()

	fmt.Printf("itest1 = %#v\n", itest1)
	fmt.Printf("itest2 = %#v\n", itest2)

	itest1.A()
	itest2.A()

	fmt.Printf("itest1 = %#v\n", itest1)
	fmt.Printf("itest2 = %#v\n", itest2)

	exOpts := &ExtendOpts{}
	exOpts.AddNode(
		&Node{Name: "x1", Value: "100"},
		&Node{Name: "x2", Value: "101"},
		&Node{Name: "x3", Value: "103"},
	)

	bookByte, err = json.Marshal(exOpts)
	if err != nil {
		fmt.Printf("err = %#v\n", err)
	}
	fmt.Printf("json(exOpts) = %s\n", string(bookByte))
	fmt.Printf("exOpts = %#v\n", exOpts)

}
