package application

import "script-go/src/application/applet/demo"

type Application struct {
}

func application() *Application {
	return &Application{}
}

func (a *Application) Run() {
	demo.Run()
}
