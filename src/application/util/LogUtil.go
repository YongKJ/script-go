package util

import (
	"fmt"
	"script-go/src/application/config"
	"script-go/src/application/pojo/dto"
)

var global = config.OfGlobal()

type LogUtil struct {
}

func newLogUtil() *LogUtil {
	return &LogUtil{}
}

func OfLogUtil() *LogUtil {
	return newLogUtil()
}

func (l *LogUtil) LoggerLine(log *dto.Log) {
	if !global.LogEnable() {
		return
	}
	fmt.Print("[" + log.ClassName() + "] " + log.MethodName() + " -> " + log.ParamName() + ": ")
	fmt.Println(log.Value())
}

func (l *LogUtil) Logger(log *dto.Log) {
	if !global.LogEnable() {
		return
	}
	fmt.Print("[" + log.ClassName() + "] " + log.MethodName() + " -> " + log.ParamName() + ": ")
	fmt.Print(log.Value())
}
