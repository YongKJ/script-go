package application

import "script-go/src/application/applet/Demo"

type Application struct {
}

func NewApplication() *Application {
	return &Application{}
}

func (a *Application) Run() {
	Demo.Run()
}
