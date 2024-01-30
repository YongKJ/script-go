package main

import "script-go/src/application"

type ScriptGo struct {
}

func scriptGo() *ScriptGo {
	return &ScriptGo{}
}

func (s *ScriptGo) run() {
	//app := new(application.Application)
	//app.Run()

	appTest := new(application.ApplicationTest)
	appTest.Test()
}

func main() {
	scriptGo := scriptGo()
	scriptGo.run()
}
