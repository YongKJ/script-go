package CHello

/*
#include <stdio.h>

void test() {
	printf("Hello from C!\n");
}
*/
import "C"
import (
	"script-go/src/application/pojo/dto/Log"
	"script-go/src/application/util/GenUtil"
	"script-go/src/application/util/LogUtil"
)

type CHello struct {
	msg string
}

func newCHello() *CHello {
	return &CHello{
		msg: GenUtil.GetValue("msg").(string),
	}
}

func (c *CHello) test() {
	C.test()
	LogUtil.LoggerLine(Log.Of("CHello", "test", "c.msg", c.msg))
}

func Run() {
	cHello := newCHello()
	cHello.test()
}
