package main

import "script-go/src/application"

type ScriptGo struct {
}

func newScriptGo() *ScriptGo {
	return &ScriptGo{}
}

func (s *ScriptGo) run() {
	//app := application.NewApplication()
	//app.Run()

	appTest := application.NewApplicationTest()
	appTest.Test()
}

func main() {
	scriptGo := newScriptGo()
	scriptGo.run()
}
