package application

import "script-go/src/application/applet/demo"

type ApplicationTest struct {
}

func NewApplicationTest() *ApplicationTest {
	return &ApplicationTest{}
}

func (a *ApplicationTest) Test() {
	demo.Run()
}
