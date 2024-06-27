package LogUtil

import (
	"fmt"
	"script-go/src/application/config"
	"script-go/src/application/pojo/dto/Log"
)

func LoggerLine(log *Log.Log) {
	if !Global.LogEnable {
		return
	}
	fmt.Print("[" + log.ClassName() + "] " + log.MethodName() + " -> " + log.ParamName() + ": ")
	fmt.Println(log.Value())
}

func Logger(log *Log.Log) {
	if !Global.LogEnable {
		return
	}
	fmt.Print("[" + log.ClassName() + "] " + log.MethodName() + " -> " + log.ParamName() + ": ")
	fmt.Print(log.Value())
}
