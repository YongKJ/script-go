package application

import "script-go/src/application/applet/Demo"

type ApplicationTest struct {
}

func NewApplicationTest() *ApplicationTest {
	return &ApplicationTest{}
}

func (a *ApplicationTest) Test() {
	Demo.Run()
}
