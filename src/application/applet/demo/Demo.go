package demo

import (
	"script-go/src/application/pojo/dto"
	"script-go/src/application/util"
)

var logUtil = util.NewLogUtil()

type Demo struct {
	msg string
}

func NewDemo() *Demo {
	return &Demo{
		"Hello world!",
	}
}

func (d *Demo) test() {
	//fmt.Println(d.msg)
	demoLogLine("test", "d.msg", d.msg)
}

func demoLogLine(methodName string, paramName string, value interface{}) {
	logUtil.LoggerLine(dto.NewLog(
		"Demo", methodName, paramName, value,
	))
}

func Run() {
	demo := NewDemo()
	demo.test()
}
