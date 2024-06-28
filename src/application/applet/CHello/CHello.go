package CHello

/*
#include <stdio.h>

void test() {
	printf("Hello from C!\n");
}
*/
import "C"

type CHello struct {
}

func newCHello() *CHello {
	return &CHello{}
}

func (c *CHello) test() {
	C.test()
}

func Run() {
	cHello := newCHello()
	cHello.test()
}
