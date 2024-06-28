package application

import "script-go/src/application/applet/CHello"

type ApplicationTest struct {
}

func NewApplicationTest() *ApplicationTest {
	return &ApplicationTest{}
}

func (a *ApplicationTest) Test() {
	CHello.Run()
	//Demo.Run()
}
