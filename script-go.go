package main

import (
	"os"
	"script-go/src/application"
)

type ScriptGo struct {
}

func newScriptGo() *ScriptGo {
	return &ScriptGo{}
}

func (s *ScriptGo) run() {
	app := application.NewApplication()
	app.Run()
}

func (s *ScriptGo) runTest() {
	appTest := application.NewApplicationTest()
	appTest.Test()
}

func main() {
	scriptGo := newScriptGo()

	args := os.Args
	if len(args) == 1 {
		scriptGo.run()
		return
	}

	flag := args[1]
	if flag == "test" {
		scriptGo.runTest()
	}
}
