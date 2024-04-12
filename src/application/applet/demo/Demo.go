package demo

import (
	"script-go/src/application/pojo/dto"
	"script-go/src/application/util"
)

var logUtil = util.OfLogUtil()

type Demo struct {
	msg string
}

func newDemo() *Demo {
	return &Demo{
		"Hello world!",
	}
}

func (d *Demo) test() {
	//fmt.Println(d.msg)
	logUtil.LoggerLine(dto.OfLog("Demo", "test", "d.msg", d.msg))
}

func Run() {
	demo := newDemo()
	demo.test()
}
