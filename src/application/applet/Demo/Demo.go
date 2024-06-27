package Demo

import (
	"script-go/src/application/pojo/dto/Log"
	"script-go/src/application/util/LogUtil"
)

type Demo struct {
}

func newDemo() *Demo {
	return &Demo{}
}

func (d *Demo) test() {
	LogUtil.LoggerLine(Log.Of("Demo", "test", "msg", "Hello world!"))
}

func Run() {
	demo := newDemo()
	demo.test()
}
