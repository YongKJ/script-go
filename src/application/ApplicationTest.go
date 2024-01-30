package application

import "script-go/src/application/applet/demo"

type ApplicationTest struct {
}

func applicationTest() *ApplicationTest {
	return &ApplicationTest{}
}

func (a *ApplicationTest) Test() {
	demo.Run()
}
