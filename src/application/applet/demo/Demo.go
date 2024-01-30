package demo

import "fmt"

type Demo struct {
	msg string
}

func demo() *Demo {
	return &Demo{
		"Hello world!",
	}
}

func (d *Demo) test() {
	fmt.Println(d.msg)
}

func Run() {
	demo := demo()
	demo.test()
}
