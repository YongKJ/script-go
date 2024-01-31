package main

import "script-go/src/application"

type ScriptGo struct {
}

func NewScriptGo() *ScriptGo {
	return &ScriptGo{}
}

func (s *ScriptGo) run() {
	//app := application.NewApplication()
	//app.Run()

	appTest := application.NewApplicationTest()
	appTest.Test()
}

func main() {
	scriptGo := NewScriptGo()
	scriptGo.run()
}
